package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const (
	musicSourceUrl = "https://music.youtube.com/"
	ytSourceUrl    = "https://www.youtube.com/"
)

// Context contains a state used for interacting with the internal YouTube api simulating a normal browser
type Context struct {
	Type                ContextType
	downloadCtxFmt      string
	searchCtxFmt        string
	key                 string
	internalUri         string
	clientVer           string
	clientName          string
	clientPlaybackNonce string
	requestNumber       int
}

// ContextType is used to specify which api to use either Youtube or YoutubeMusic
type ContextType int

const (
	Youtube      ContextType = 0
	YoutubeMusic ContextType = 1
)

// NewContext creates a new Context given a ContextType specifying what api to use
func NewContext(contextType ContextType) (Context, error) {
	ctx := Context{Type: contextType, clientPlaybackNonce: GenerateClientPlaybackNonce(), requestNumber: 1}
	if contextType == YoutubeMusic {
		reqUrl, _ := url.Parse(musicSourceUrl)
		req := &http.Request{
			Method: "GET",
			URL:    reqUrl,
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
			},
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return Context{}, err
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return Context{}, err
		}
		musicCfg := regexp.MustCompile(`ytcfg.set\(({"CLIENT_CANARY_STATE".+})`).FindSubmatch(body)
		if len(musicCfg) <= 1 {
			return Context{}, errors.New("youtube music config set not found in response")
		}
		cfgSet := musicConfigSet{}
		if err = json.NewDecoder(bytes.NewReader(musicCfg[1])).Decode(&cfgSet); err != nil {
			return Context{}, err
		}
		ctx.internalUri = "https://music.youtube.com/youtubei/v1"
		ctx.clientName = cfgSet.InnerTubeClientName
		ctx.clientVer = cfgSet.InnerTubeClientVersion
		ctx.key = cfgSet.InnerTubeApiKey
		ctx.downloadCtxFmt = fmt.Sprintf(`{
            "videoId": "%s",
            "context": {
            "client": {
                "hl": "en",
                "gl": "US",
                "clientName": "%s",
                "clientVersion": "%s"
            }
            },
            "playbackContext": {
            "contentPlaybackContext": {
                "referer": "https://music.youtube.com/watch?v=%s"
            }
            },
            "cpn": "%s"
        }`,
			"%s",
			ctx.clientName,
			ctx.clientVer,
			"%s",
			ctx.clientPlaybackNonce,
		)
		ctx.searchCtxFmt = fmt.Sprintf(`{
            "context": {
                "client": {
                "hl": "en",
                "gl": "US",
                "clientName": "%s",
                "clientVersion": "%s",
            },
            },
            "query": "%s"
        }`,
			ctx.clientName,
			ctx.clientVer,
			"%s",
		)
	} else if contextType == Youtube {
		res, err := http.Get(ytSourceUrl)
		if err != nil {
			return Context{}, err
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return Context{}, err
		}
		musicCfg := regexp.MustCompile(`ytcfg.set\(({"CLIENT_CANARY_STATE".+})`).FindSubmatch(body)
		if len(musicCfg) <= 1 {
			return Context{}, errors.New("youtube music config set not found in response")
		}
		cfgSet := youtubeConfigSet{}
		if err = json.NewDecoder(bytes.NewReader(musicCfg[1])).Decode(&cfgSet); err != nil {
			return Context{}, err
		}
		ctx.internalUri = "https://www.youtube.com/youtubei/v1"
		ctx.clientName = cfgSet.InnerTubeClientName
		ctx.clientVer = cfgSet.InnerTubeClientVersion
		ctx.key = cfgSet.InnerTubeApiKey
		ctx.downloadCtxFmt = fmt.Sprintf(`{
            "videoId": "%s",
            "context": {
            "client": {
                "hl": "en",
                "gl": "US",
                "clientName": "%s",
                "clientVersion": "%s",
            }
            },
            "playbackContext": {
            "contentPlaybackContext": {
                "referer": "https://youtube.com/watch?v=%s"
            }
            },
            "cpn": "%s"
        }`,
			"%s",
			ctx.clientName,
			ctx.clientVer,
			"%s",
			ctx.clientPlaybackNonce,
		)
		ctx.searchCtxFmt = fmt.Sprintf(`{
            "context": {
                "client": {
                "hl": "en",
                "gl": "US",
                "clientName": "%s",
                "clientVersion": "%s",
            },
            },
            "query": "%s"
        }`,
			ctx.clientName,
			ctx.clientVer,
			"%s",
		)
	}
	return ctx, nil
}

func (c Context) String() string {
	return fmt.Sprintf("cver: %s, cname: %s, key: %s, url: %s", c.clientVer, c.clientName, c.key, c.internalUri)
}
