package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type (
	// QueuesService is a service to deal with queues.
	QueuesService struct {
		client *Client
	}

	// QueuesListResp describes queues method response.
	QueuesListResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Queues []QueueElement `json:"queues"`
		} `json:"result"`
	}

	// QueuesGetResp describes queue method response
	QueuesGetResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Tracks []struct {
				TrackID string `json:"trackId"`
				AlbumID string `json:"albumId"`
				From    string `json:"from"`
			} `json:"tracks"`
			CurrentIndex int       `json:"currentIndex"`
			Modified     time.Time `json:"modified"`
		} `json:"result"`
	}

	QueueElement struct {
		ID       string    `json:"id"`
		Modified time.Time `json:"modified"`
	}
)

// List returns list of existed queues.
func (s *QueuesService) List(
	ctx context.Context,
) (*QueuesListResp, *http.Response, error) {

	req, err := s.client.NewRequest(http.MethodGet, "queues", nil)
	if err != nil {
		return nil, nil, err
	}

	queues := new(QueuesListResp)
	resp, err := s.client.Do(ctx, req, queues)
	return queues, resp, err
}

// Get returns playlist of the user by kind
func (s *QueuesService) Get(
	ctx context.Context,
	queueID string,
) (*QueuesGetResp, *http.Response, error) {
	uri := fmt.Sprintf("queues/%s", queueID)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	queue := new(QueuesGetResp)
	resp, err := s.client.Do(ctx, req, queue)
	return queue, resp, err
}
