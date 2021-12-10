package spotify

import "time"

type AccessTokenResponse struct {
	ClientId                         string `json:"clientId"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs int64  `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      bool   `json:"isAnonymous"`
}

type SpotifyPlaylist struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		Url    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Name  string `json:"name"`
	Owner struct {
		DisplayName  string `json:"display_name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id   string `json:"id"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"owner"`
	PrimaryColor interface{} `json:"primary_color"`
	Public       bool        `json:"public"`
	SnapshotId   string      `json:"snapshot_id"`
	Tracks       struct {
		Href  string `json:"href"`
		Items []struct {
			AddedAt time.Time `json:"added_at"`
			AddedBy struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"added_by"`
			IsLocal      bool        `json:"is_local"`
			PrimaryColor interface{} `json:"primary_color"`
			Track        struct {
				Album struct {
					AlbumType *string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						Id   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						Uri  string `json:"uri"`
					} `json:"artists"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalUrls     struct {
						Spotify string `json:"spotify,omitempty"`
					} `json:"external_urls"`
					Href   *string `json:"href"`
					Id     *string `json:"id"`
					Images []struct {
						Height int    `json:"height"`
						Url    string `json:"url"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string  `json:"name"`
					ReleaseDate          *string `json:"release_date"`
					ReleaseDatePrecision *string `json:"release_date_precision"`
					TotalTracks          int     `json:"total_tracks,omitempty"`
					Type                 string  `json:"type"`
					Uri                  *string `json:"uri"`
				} `json:"album"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify,omitempty"`
					} `json:"external_urls"`
					Href *string `json:"href"`
					Id   *string `json:"id"`
					Name string  `json:"name"`
					Type string  `json:"type"`
					Uri  *string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMs       int      `json:"duration_ms"`
				Episode          bool     `json:"episode,omitempty"`
				Explicit         bool     `json:"explicit"`
				ExternalIds      struct {
					Isrc string `json:"isrc,omitempty"`
				} `json:"external_ids"`
				ExternalUrls struct {
					Spotify string `json:"spotify,omitempty"`
				} `json:"external_urls"`
				Href        string  `json:"href"`
				Id          string  `json:"id"`
				IsLocal     bool    `json:"is_local"`
				Name        string  `json:"name"`
				Popularity  int     `json:"popularity"`
				PreviewUrl  *string `json:"preview_url"`
				Track       bool    `json:"track,omitempty"`
				TrackNumber int     `json:"track_number"`
				Type        string  `json:"type"`
				Uri         string  `json:"uri"`
			} `json:"track"`
			VideoThumbnail struct {
				Url interface{} `json:"url"`
			} `json:"video_thumbnail"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type SpotifyTracks struct {
	Href  string `json:"href"`
	Items []struct {
		AddedAt time.Time `json:"added_at"`
		AddedBy struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"added_by"`
		IsLocal      bool        `json:"is_local"`
		PrimaryColor interface{} `json:"primary_color"`
		Track        struct {
			Album struct {
				AlbumType string `json:"album_type"`
				Artists   []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					Id   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Uri  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				Id     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				TotalTracks          int    `json:"total_tracks"`
				Type                 string `json:"type"`
				Uri                  string `json:"uri"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Episode          bool     `json:"episode"`
			Explicit         bool     `json:"explicit"`
			ExternalIds      struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string  `json:"href"`
			Id          string  `json:"id"`
			IsLocal     bool    `json:"is_local"`
			Name        string  `json:"name"`
			Popularity  int     `json:"popularity"`
			PreviewUrl  *string `json:"preview_url"`
			Track       bool    `json:"track"`
			TrackNumber int     `json:"track_number"`
			Type        string  `json:"type"`
			Uri         string  `json:"uri"`
		} `json:"track"`
		VideoThumbnail struct {
			Url interface{} `json:"url"`
		} `json:"video_thumbnail"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}
