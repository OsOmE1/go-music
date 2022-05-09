package youtube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ParseVideo returns a Result given a YouTube video id
func ParseVideo(id string) (Result, error) {
	res, err := http.Get(fmt.Sprintf("https://youtube.com/watch?v=%v", id))
	if err != nil {
		return Result{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Result{}, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Result{}, err
	}
	rawData := regexp.MustCompile(`<script.*>var ytInitialPlayerResponse = ({.+});.*<\/script>`).FindSubmatch(resBody)[1]
	data := videoAPIResponse{}
	if err = json.Unmarshal(rawData, &data); err != nil {
		return Result{}, err
	}
	result := Result{
		Type:       TypeVideo,
		ID:         data.VideoDetails.VideoId,
		Thumbnails: []string{},
		Title:      data.VideoDetails.Title,
	}
	for _, thumbnail := range data.VideoDetails.Thumbnail.Thumbnails {
		result.Thumbnails = append(result.Thumbnails, thumbnail.Url)
	}
	result.Length, err = strconv.Atoi(data.VideoDetails.LengthSeconds)
	return result, nil
}

// ParseMusicTrack returns a Result given a YouTube Music track id
func (c Context) ParseMusicTrack(id string) (Result, error) {
	reqBody := bytes.NewReader([]byte(fmt.Sprintf(c.downloadCtxFmt, id, c.clientPlaybackNonce)))
	reqUrl, _ := url.Parse(fmt.Sprintf("%s/player?key=%s&prettyPrint=false", c.internalUri, c.key))
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"User-Agent":               {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
			"x-youtube-client-version": {c.clientVer},
			"Referer":                  {fmt.Sprintf("https://music.youtube.com/watch?v=%s", id)},
		},
		Body: io.NopCloser(reqBody),
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Result{}, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Result{}, err
	}
	data := musicPlayerParseResponse{}
	if err = json.Unmarshal(resBody, &data); err != nil {
		return Result{}, err
	}
	result := Result{
		Type:       TypeTrack,
		ID:         data.VideoDetails.VideoId,
		Thumbnails: []string{},
		Title:      data.VideoDetails.Title,
	}
	for _, thumbnail := range data.VideoDetails.Thumbnail.Thumbnails {
		result.Thumbnails = append(result.Thumbnails, thumbnail.Url)
	}
	result.Length, err = strconv.Atoi(data.VideoDetails.LengthSeconds)
	return result, nil
}

// ParsePlaylist returns a slice of Results given a YouTube playlist id
func ParsePlaylist(id string) ([]Result, error) {
	res, err := http.Get(fmt.Sprintf("https://youtube.com/playlist?list=%v", id))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	rawData := regexp.MustCompile(`<script.*>var ytInitialData = ({.+});.*<\/script>`).FindSubmatch(resBody)[1]
	data := playlistAPIResponse{}
	if err = json.Unmarshal(rawData, &data); err != nil {
		return nil, err
	}
	var results []Result
	if len(data.Contents.TwoColumnBrowseResultsRenderer.Tabs) == 0 {
		return nil, nil
	}
	if len(data.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents) == 0 {
		return nil, nil
	}
	if len(data.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents) == 0 {
		return nil, nil
	}
	for _, item := range data.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].PlaylistVideoListRenderer.Contents {
		if item.PlaylistVideoRenderer.VideoId == "" {
			continue
		}
		title := ""
		if len(item.PlaylistVideoRenderer.Title.Runs) != 0 {
			title = item.PlaylistVideoRenderer.Title.Runs[0].Text
		}
		results = append(results, parseVideoRenderer(videoRenderer{
			VideoId: item.PlaylistVideoRenderer.VideoId,
			Title: struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
			}{
				Runs: []struct {
					Text string `json:"text"`
				}{
					{title},
				},
			},
			LengthText: item.PlaylistVideoRenderer.LengthText,
		}))
	}
	return results, nil
}
