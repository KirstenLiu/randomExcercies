package main

import "time"

type Playlist struct{
	Kind string `json:"kind"`
	PlaylistID string `json:"playlistId"`
}

type Channel struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode string `json:"regionCode"`
	PageInfo struct {
		TotalResults int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID struct {
			Kind string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id,omitempty"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID string `json:"channelId"`
			Title string `json:"title"`
			Description string `json:"description"`
			Thumbnails struct {
				Default struct {
					URL string `json:"url"`
					Width int `json:"width"`
					Height int `json:"height"`
				} `json:"default"`
				Medium struct {
					URL string `json:"url"`
					Width int `json:"width"`
					Height int `json:"height"`
				} `json:"medium"`
				High struct {
					URL string `json:"url"`
					Width int `json:"width"`
					Height int `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
			PublishTime time.Time `json:"publishTime"`
		} `json:"snippet"`
		p []Playlist 
	} `json:"items"`
}

type ChannelList struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			Localized struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
			Country string `json:"country"`
		} `json:"snippet"`
		ContentDetails struct {
			RelatedPlaylists struct {
				Likes   string `json:"likes"`
				Uploads string `json:"uploads"`
			} `json:"relatedPlaylists"`
		} `json:"contentDetails"`
	} `json:"items"`
}

type Videos struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind           string `json:"kind"`
		Etag           string `json:"etag"`
		ID             string `json:"id"`
		ContentDetails struct {
			Duration        string `json:"duration"`
			Dimension       string `json:"dimension"`
			Definition      string `json:"definition"`
			Caption         string `json:"caption"`
			LicensedContent bool   `json:"licensedContent"`
			ContentRating   struct {
			} `json:"contentRating"`
			Projection string `json:"projection"`
		} `json:"contentDetails"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type PlaylistItem struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	Items         []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle string `json:"channelTitle"`
			PlaylistID   string `json:"playlistId"`
			Position     int    `json:"position"`
			ResourceID   struct {
				Kind    string `json:"kind"`
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
			VideoOwnerChannelTitle string `json:"videoOwnerChannelTitle"`
			VideoOwnerChannelID    string `json:"videoOwnerChannelId"`
		} `json:"snippet"`
		ContentDetails struct {
			VideoID          string    `json:"videoId"`
			VideoPublishedAt time.Time `json:"videoPublishedAt"`
		} `json:"contentDetails"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}