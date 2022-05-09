package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OsOmE1/go-music/youtube"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

const (
	// Either one of the artists, the title or the length has to match
	matchThreshold int = 1
	// The margin for the duration of a track to match
	lengthMargin int = 2
)

// Match contains a matching pair of a youtube.Result and a spotify Track
type Match struct {
	Spotify Track
	YTMusic youtube.Result
}

// Track contains spotify track
type Track struct {
	ID      string
	Name    string
	Artists []string
	Length  int
}

// Playlist contains data on a specific spotify playlist
type Playlist struct {
	ID     string
	Name   string
	Tracks []Track
}

// Matches pass nil pointer to disable cache usage
func (p Playlist) Matches(cache *Cache, ctx youtube.Context) ([]Match, error) {
	var results []Match
	resultLock := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(len(p.Tracks))
	for _, t := range p.Tracks {
		track := t
		go func() {
			defer wg.Done()

			result, err := track.YoutubeMatch(cache, ctx)
			if err != nil {
				return
			}
			resultLock.Lock()
			results = append(results, result)
			resultLock.Unlock()
		}()
	}

	wg.Wait()
	return results, nil
}

func (t Track) YoutubeMatch(cache *Cache, ctx youtube.Context) (Match, error) {
	if cache != nil {
		r := cache.Result(t.ID)
		if r != nil {
			return Match{Spotify: t, YTMusic: *r}, nil
		}
	}
	results, err := ctx.Search(t.queryName())
	if err != nil {
		return Match{}, err
	}
	match := Match{Spotify: t}
	score := 0
	for _, result := range results {
		localScore := 0
		if result.Title == t.Name {
			localScore++
		}
		for _, artist := range t.Artists {
			for _, resultArtist := range result.Artists {
				if strings.TrimSpace(strings.ToLower(artist)) == strings.TrimSpace(strings.ToLower(resultArtist)) {
					localScore++
				}
			}
		}
		if t.Length+lengthMargin >= result.Length && result.Length >= t.Length-lengthMargin {
			localScore++
		}
		if localScore > score {
			match.YTMusic = result
			score = localScore
		}
	}
	if score == -1 {
		return match, errors.New("no match found")
	}
	if score < matchThreshold {
		return match, errors.New("match found but below match threshold")
	}

	err = cache.AddToCache(t.ID, match.YTMusic)
	if err != nil {
		return match, err
	}
	return match, nil
}

func ParsePlaylist(shareUrl string) (Playlist, error) {
	playlist := Playlist{}
	matches := regexp.MustCompile(`https:\/\/open.spotify.com\/playlist\/([A-Za-z0-9]+)`).FindStringSubmatch(shareUrl)
	if len(matches) <= 1 {
		return playlist, errors.New("invalid share url, should be in the form of: https://open.spotify.com/playlist/<id>")
	}
	apiKey, err := getInternalApiKey()
	if err != nil {
		return playlist, err
	}
	reqUrl, _ := url.Parse(fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", matches[1]))
	req := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{
			"Authorization": {fmt.Sprintf("Bearer %s", apiKey)},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return playlist, fmt.Errorf("error while retrieving %v: %v", reqUrl, err)
	}

	if res.StatusCode != http.StatusOK {
		return playlist, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return playlist, err
	}
	spotifyPlaylist := SpotifyPlaylist{}
	if err = json.Unmarshal(resBody, &spotifyPlaylist); err != nil {
		return playlist, err
	}
	playlist.ID = spotifyPlaylist.Id
	playlist.Name = spotifyPlaylist.Name
	playlist.Tracks = []Track{}
	for _, item := range spotifyPlaylist.Tracks.Items {
		t := Track{
			ID:      item.Track.Id,
			Name:    item.Track.Name,
			Artists: []string{},
			Length:  item.Track.DurationMs / 1000,
		}
		for _, artist := range item.Track.Artists {
			t.Artists = append(t.Artists, artist.Name)
		}
		playlist.Tracks = append(playlist.Tracks, t)
	}

	next := spotifyPlaylist.Tracks.Next
	for next != "" {
		reqUrl, _ := url.Parse(next)
		req := &http.Request{
			Method: "GET",
			URL:    reqUrl,
			Header: map[string][]string{
				"Authorization": {fmt.Sprintf("Bearer %s", apiKey)},
			},
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return Playlist{}, err
		}
		if res.StatusCode != http.StatusOK {
			return Playlist{}, fmt.Errorf("unexpected status code: %v", res.StatusCode)
		}
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return Playlist{}, err
		}
		trackData := SpotifyTracks{}
		if err = json.Unmarshal(resBody, &trackData); err != nil {
			return Playlist{}, err
		}
		for _, item := range trackData.Items {
			t := Track{
				ID:      item.Track.Id,
				Name:    item.Track.Name,
				Artists: []string{},
			}
			for _, artist := range item.Track.Artists {
				t.Artists = append(t.Artists, artist.Name)
			}
			playlist.Tracks = append(playlist.Tracks, t)
		}
		next = trackData.Next
	}
	return playlist, nil
}

func (t Track) queryName() string {
	return fmt.Sprintf("%s %s", t.Name, strings.Join(t.Artists, ", "))
}
