package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

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
	data := VideoAPIResponse{}
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
	data := PlaylistAPIResponse{}
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
		results = append(results, parseVideoRenderer(VideoRenderer{
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
