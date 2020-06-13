package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	// TracksService is a service to deal with tracks.
	TracksService struct {
		client *Client
	}

	TracksDownloadInfoRsp struct {
		InvocationInfo InvocationInfo        `json:"invocationInfo"`
		Error          Error                 `json:"error"`
		Result         []DownloadInfoElement `json:"result"`
	}

	DownloadInfoElement struct {
		Codec   string `json:"codec"`
		Bitrate int    `json:"bitrateInKbps"`
		URL     string `json:"downloadInfoUrl"`
	}

	// TracksGetMultipleRsp describes tracks method response.
	TracksGetMultipleRsp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         []TrackElement `json:"result"`
	}

	TrackElement struct {
		ID             string   `json:"id"`
		DurationMs     int      `json:"durationMs"`
		Available      bool     `json:"available"`
		AvailableAsRbt bool     `json:"availableAsRbt"`
		Explicit       bool     `json:"explicit"`
		StorageDir     string   `json:"storageDir"`
		Title          string   `json:"title"`
		Version        string   `json:"version,omitempty"`
		Regions        []string `json:"regions"`
		Albums         []struct {
			ID                  int           `json:"id"`
			StorageDir          string        `json:"storageDir"`
			OriginalReleaseYear int           `json:"originalReleaseYear"`
			Year                int           `json:"year"`
			Title               string        `json:"title"`
			Artists             []interface{} `json:"artists"`
			CoverURI            string        `json:"coverUri"`
			TrackCount          int           `json:"trackCount"`
			Genre               string        `json:"genre"`
			Available           bool          `json:"available"`
			TrackPosition       struct {
				Volume int `json:"volume"`
				Index  int `json:"index"`
			} `json:"trackPosition"`
		} `json:"albums"`
		Artists []struct {
			ID    int `json:"id"`
			Cover struct {
				Type   string `json:"type"`
				Prefix string `json:"prefix"`
				URI    string `json:"uri"`
			} `json:"cover"`
			Composer   bool          `json:"composer"`
			Various    bool          `json:"various"`
			Name       string        `json:"name"`
			Decomposed []interface{} `json:"decomposed"`
		} `json:"artists"`
		CoverURI string `json:"coverUri"`
		Type     string `json:"type"`
	}
)

// GetTracks returns multiple tracks
func (s *TracksService) GetTracks(
	ctx context.Context,
	trackIDs []string,
) (*TracksGetMultipleRsp, *http.Response, error) {
	form := url.Values{}
	form.Set("track_ids", strings.Join(trackIDs, ","))

	req, err := s.client.NewRequest(http.MethodPost, "tracks", form)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tracks := new(TracksGetMultipleRsp)
	resp, err := s.client.Do(ctx, req, tracks)
	return tracks, resp, err
}

func (s *TracksService) GetDownloadInfo(
	ctx context.Context,
	trackID string,
) (*TracksDownloadInfoRsp, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tracks/%s/download-info", trackID), nil)
	if err != nil {
		return nil, nil, err
	}

	download := new(TracksDownloadInfoRsp)
	resp, err := s.client.Do(ctx, req, download)
	return download, resp, err
}
