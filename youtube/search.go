package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type YoutubeType int

const (
	TypeVideo    YoutubeType = 0
	TypePlaylist YoutubeType = 1
	TypeTrack    YoutubeType = 2
)

// Result contains data from a youtube search or id query
type Result struct {
	Type       YoutubeType `json:"type"`
	ID         string      `json:"id"`
	Thumbnails []string    `json:"thumbnails,omitempty"`
	Title      string      `json:"title"`
	Length     int         `json:"length"`
	Owner      Channel     `json:"owner,omitempty"`
	Artists    []string    `json:"artists"`
	Children   []Result    `json:"children,omitempty"`
}

type Channel struct {
	Name     string `json:"name,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	Icon     string `json:"icon,omitempty"`
}

// Search uses a query to find a youtube video or track
func (c Context) Search(query string) ([]Result, error) {
	if c.Type == Youtube {
		return c.searchYoutube(query)
	} else if c.Type == YoutubeMusic {
		return c.searchMusic(query)
	}
	return nil, errors.New("invalid contextType")
}

func (c Context) searchMusic(query string) ([]Result, error) {
	reqBody := bytes.NewReader([]byte(fmt.Sprintf(c.searchCtxFmt, query)))
	reqUrl, _ := url.Parse(fmt.Sprintf("%s/search?key=%s", c.internalUri, c.key))
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"x-youtube-client-version": {c.clientVer},
			"Referer":                  {fmt.Sprintf("%ssearch?q=%v", musicSourceUrl, query)},
		},
		Body: io.NopCloser(reqBody),
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	c.requestNumber++
	if res.StatusCode != http.StatusOK {
		resBody, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(resBody))
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody := musicSearchResponse{}
	if err = json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}
	var results []Result
	for _, renderer := range resBody.Contents.TabbedSearchResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents {
		for _, item := range renderer.MusicShelfRenderer.Contents {
			if item.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId != "" {
				results = append(results, parseMusicItemRenderer(item.MusicResponsiveListItemRenderer))
			}
		}
	}
	return results, nil
}

func (c Context) searchYoutube(query string) ([]Result, error) {
	reqBody := bytes.NewReader([]byte(fmt.Sprintf(c.searchCtxFmt, query)))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/search?key=%s", c.internalUri, c.key), reqBody)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	c.requestNumber++

	resBody := youtubeSearchResponse{}
	if err = json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}
	var results []Result
	for _, renderer := range resBody.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, item := range renderer.ItemSectionRenderer.Contents {
			if item.VideoRenderer.VideoId != "" {
				results = append(results, parseVideoRenderer(item.VideoRenderer))
			} else if item.PlaylistRenderer.PlaylistId != "" {
				result := Result{
					Type:     TypePlaylist,
					ID:       item.PlaylistRenderer.PlaylistId,
					Title:    item.PlaylistRenderer.Title.SimpleText,
					Children: []Result{},
				}
				for _, thumbnails := range item.PlaylistRenderer.Thumbnails {
					for _, thumbnail := range thumbnails.Thumbnails {
						result.Thumbnails = append(result.Thumbnails, thumbnail.Url)
					}
				}
				result.Length, _ = strconv.Atoi(item.PlaylistRenderer.VideoCount)
				for _, video := range item.PlaylistRenderer.Videos {
					result.Children = append(result.Children, parseVideoRenderer(videoRenderer{
						VideoId: video.ChildVideoRenderer.VideoId,
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
								{video.ChildVideoRenderer.Title.SimpleText},
							},
						},
						LengthText: video.ChildVideoRenderer.LengthText,
					}))
				}
				results = append(results, result)
			}
		}
	}
	return results, nil
}

func parseMusicItemRenderer(v musicItemRenderer) Result {
	result := Result{
		Type:  TypeTrack,
		ID:    v.PlaylistItemData.VideoId,
		Title: v.FlexColumns[0].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[0].Text,
	}
	duration := ""
	result.Artists = []string{}
	if len(v.FlexColumns) > 1 {
		for _, run := range v.FlexColumns[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs {
			if run.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType == "MUSIC_PAGE_TYPE_ARTIST" {
				artist := strings.Replace(run.Text, "\u200b", "", 1)
				result.Artists = append(result.Artists, artist)
			}
			if strings.Contains(run.Text, ":") {
				duration = run.Text
			}
		}
	}

	for _, thumbnail := range v.Thumbnail.MusicThumbnailRenderer.Thumbnail.Thumbnails {
		result.Thumbnails = append(result.Thumbnails, thumbnail.Url)
	}

	if duration == "" {
		return result
	}
	durationStrings := strings.Split(duration, ":")
	totalDuration := time.Duration(0)
	unit := time.Second
	for i := len(durationStrings) - 1; i >= 0; i-- {
		duration, err := strconv.Atoi(durationStrings[i])
		if err != nil {
			return result // If we can't parse the duration, just omit it for now.
		}
		totalDuration += time.Duration(duration) * unit
		if unit == time.Second {
			unit = time.Minute
			continue
		}
		if unit == time.Minute {
			unit = time.Hour
		}
	}
	result.Length = int(totalDuration.Seconds())
	return result
}

func parseVideoRenderer(v videoRenderer) Result {
	result := Result{
		Type:  TypeVideo,
		ID:    v.VideoId,
		Title: v.Title.Runs[0].Text,
	}
	if len(v.OwnerText.Runs) != 0 {
		result.Owner.Name = v.OwnerText.Runs[0].Text
	}
	if len(v.ChannelThumbnailSupportedRenderers.ChannelThumbnailWithLinkRenderer.Thumbnail.Thumbnails) != 0 {
		result.Owner.Icon = v.ChannelThumbnailSupportedRenderers.ChannelThumbnailWithLinkRenderer.Thumbnail.Thumbnails[0].Url
	}
	for _, thumbnail := range v.Thumbnail.Thumbnails {
		result.Thumbnails = append(result.Thumbnails, thumbnail.Url)
	}
	for _, badge := range v.OwnerBadges {
		if badge.MetadataBadgeRenderer.Style == "BADGE_STYLE_TYPE_VERIFIED" {
			result.Owner.Verified = true
		}
	}
	durationStrings := strings.Split(v.LengthText.SimpleText, ":")
	totalDuration := time.Duration(0)
	unit := time.Second
	for i := len(durationStrings) - 1; i >= 0; i-- {
		duration, err := strconv.Atoi(durationStrings[i])
		if err != nil {
			return result // If we can't parse the duration, just omit it for now.
		}
		totalDuration += time.Duration(duration) * unit
		if unit == time.Second {
			unit = time.Minute
		}
		if unit == time.Minute {
			unit = time.Hour
		}
	}
	result.Length = int(totalDuration.Seconds())
	return result
}
