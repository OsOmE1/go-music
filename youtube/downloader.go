package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Downloader struct {
	cpn string
	rn  int
}

type FormatType int

const (
	AudioFormat FormatType = 0
	VideoFormat FormatType = 1
)

type Format struct {
	Type FormatType
	// General
	ContentType   string
	Codec         string
	ContentLength int
	DurationMS    int

	// Audio
	AudioQuality   string
	AverageBitrate int
	SampleRate     int
	Channels       int

	// Video
	Quality string
	Fps     int
	Bitrate int
	Width   int
	Height  int

	// Sources
	url             string
	signatureCipher string
}

type Chunk struct {
	Data []byte
	High int
	Low  int
}

func (c Context) Formats(result Result) ([]Format, error) {
	switch result.Type {
	case TypeVideo:
		if c.Type != Youtube {
			return nil, errors.New("cannot get formats for youtube video with youtube music context")
		}
		return c.ytFormats(result)
	case TypeTrack:
		if c.Type != YoutubeMusic {
			return nil, errors.New("cannot get formats for youtube music with youtube context")
		}
		return c.musicFormats(result)
	default:
		return nil, fmt.Errorf("cannot get formats for type: %d", result.Type)
	}
}

func (c Context) ytFormats(result Result) ([]Format, error) {
	reqBody := bytes.NewReader([]byte(fmt.Sprintf(c.downloadCtxFmt, result.ID, c.clientPlaybackNonce)))
	reqUrl, _ := url.Parse(fmt.Sprintf("%s/player?key=%s", c.internalUri, c.key))
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"User-Agent":               {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
			"x-youtube-client-version": {c.clientVer},
			"Referer":                  {fmt.Sprintf("https://www.youtube.com/watch?v=%v", result.ID)},
		},
		Body: io.NopCloser(reqBody),
	}
	c.requestNumber++
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody := YoutubePlayerResponse{}
	if err = json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}
	return parseStreamingData(resBody.StreamingData), nil
}

func (c Context) musicFormats(result Result) ([]Format, error) {
	reqBody := bytes.NewReader([]byte(fmt.Sprintf(c.downloadCtxFmt, result.ID, c.clientPlaybackNonce)))
	reqUrl, _ := url.Parse(fmt.Sprintf("%s/player?key=%v", c.internalUri, c.key))
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"User-Agent":               {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
			"x-youtube-client-version": {c.clientVer},
			"Referer":                  {fmt.Sprintf("%swatch?v=%s", musicSourceUrl, result.ID)},
		},
		Body: io.NopCloser(reqBody),
	}
	c.requestNumber++
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody := PlayerQueryResponse{}
	if err = json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	return parseStreamingData(resBody.StreamingData), nil
}

func parseStreamingData(data StreamingData) []Format {
	var formats []Format
	for _, format := range data.AdaptiveFormats {
		formatType := AudioFormat
		if format.AudioQuality == "" {
			formatType = VideoFormat
		}
		clen, err := strconv.ParseInt(format.ContentLength, 10, 32)
		if err != nil {
			continue
		}
		dur, err := strconv.ParseInt(format.ApproxDurationMs, 10, 32)
		if err != nil {
			continue
		}
		var sample int64 = 0
		if formatType == AudioFormat {
			sample, err = strconv.ParseInt(format.AudioSampleRate, 10, 32)
			if err != nil {
				continue
			}
		}
		formats = append(formats, Format{
			Type:            formatType,
			ContentType:     strings.Split(format.MimeType, ";")[0],
			Codec:           strings.Replace(strings.Split(format.MimeType, "; codecs=")[1], "\"", "", 2),
			ContentLength:   int(clen),
			AudioQuality:    format.AudioQuality,
			AverageBitrate:  format.AverageBitrate,
			DurationMS:      int(dur),
			SampleRate:      int(sample),
			Channels:        format.AudioChannels,
			url:             format.Url,
			Fps:             format.Fps,
			Bitrate:         format.Bitrate,
			Quality:         format.Quality,
			Width:           format.Width,
			Height:          format.Height,
			signatureCipher: format.SignatureCipher,
		})
	}
	return formats
}

func (c Context) DownloadFormatChunk(format Format, low, high int) (Chunk, error) {
	cipher := strings.Split(format.signatureCipher, "&")
	downloadUrl := format.url
	sig := ""
	for _, item := range cipher {
		if strings.Contains(item, "url=") {
			tmpurl, err := url.QueryUnescape(strings.Replace(item, "url=", "", 1))
			if err != nil {
				return Chunk{}, fmt.Errorf("error unescaping url %v", err)
			}
			downloadUrl = tmpurl
		}
		if strings.Contains(item, "s=") {
			tmpsig, err := url.QueryUnescape(strings.Replace(item, "s=", "", 1))
			if err != nil {
				return Chunk{}, fmt.Errorf("error unescaping lsig %v", err)
			}
			sig = tmpsig
		}
	}
	if sig != "" {
		downloadUrl = fmt.Sprintf("%s&alr=yes&sig=%s&cpn=%s&cver=%s&rbuf=%d",
			downloadUrl, ComputeSignature(sig), c.clientPlaybackNonce, c.clientVer, 0)
	}
	if downloadUrl == "" {
		return Chunk{}, errors.New("no url found for format")
	}

	if low < 0 {
		low = 0
	}
	if high > format.ContentLength {
		high = format.ContentLength - 1
	}

	res, err := http.Get(fmt.Sprintf("%s&rn=%d&range=%d-%d", downloadUrl, c.requestNumber, low, high))
	if err != nil {
		return Chunk{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Chunk{}, fmt.Errorf("unexpected http reponse code: %v", res.StatusCode)
	}
	c.requestNumber++

	resBody, _ := ioutil.ReadAll(res.Body)
	return Chunk{
		Data: resBody,
		High: high,
		Low:  low,
	}, nil
}

func (c Context) DownloadFormatSync(format Format, maxChunkSize int) ([]byte, error) {
	chunkSize := maxChunkSize
	if chunkSize > format.ContentLength {
		chunkSize = format.ContentLength - 1
	}
	chunkAmount := int(math.Ceil(float64(format.ContentLength) / float64(chunkSize)))
	var out []byte

	for i := 0; i < chunkAmount; i++ {
		low, high := i*chunkSize, i*chunkSize+chunkSize
		if high >= format.ContentLength {
			high = format.ContentLength - 1
		}
		chunk, err := c.DownloadFormatChunk(format, low, high)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			out = append(out, chunk.Data...)
			continue
		}
		out = append(out, chunk.Data[1:]...)
	}
	return out, nil
}

func (c Context) DownloadFormatParallel(format Format, maxChunkSize int) ([]byte, error) {
	chunkSize := maxChunkSize
	if chunkSize > format.ContentLength {
		chunkSize = format.ContentLength - 1
	}
	chunkAmount := int(math.Ceil(float64(format.ContentLength) / float64(chunkSize)))
	var chunks []Chunk
	chunksLock := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(chunkAmount)
	for i := 0; i < chunkAmount; i++ {
		index := i
		go func() {
			defer wg.Done()

			low, high := index*chunkSize, index*chunkSize+chunkSize
			if high >= format.ContentLength {
				high = format.ContentLength - 1
			}
			chunk, err := c.DownloadFormatChunk(format, low, high)
			if err != nil {
				return
			}
			chunksLock.Lock()
			chunks = append(chunks, chunk)
			chunksLock.Unlock()
		}()
	}
	wg.Wait()

	sort.SliceStable(chunks, func(i, j int) bool {
		return chunks[i].Low < chunks[j].Low
	})
	var out []byte
	for i, chunk := range chunks {
		if i == 0 {
			out = append(out, chunk.Data...)
			continue
		}
		out = append(out, chunk.Data[1:]...)
	}
	return out, nil
}
