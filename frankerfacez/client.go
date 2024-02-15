package frankerfacez

import (
	"context"
	"net/http"
	"net/url"
)

type Client interface {
	Get(ctx context.Context, url url.URL) (*http.Response, error)

	ApiV1
}

type ApiV1 interface {
	// GetEmotes Get emote by some parameters
	GetEmotes(ctx context.Context, in ApiV1EmotesRequest) (*ApiV1EmotesResponse, error)

	// GetRoomByName Get room info and emote set in room
	GetRoomByName(ctx context.Context, in ApiV1RoomByNameRequest) (*ApiV1RoomByNameResponse, error)
}

type client struct {
	client http.Client
}

func NewClient() Client {
	return &client{
		client: http.Client{},
	}
}

func (c *client) Get(ctx context.Context, url url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	errChan := make(chan error, 1)
	resChan := make(chan *http.Response, 1)

	go func() {
		res, err := c.client.Do(req)
		if err != nil {
			errChan <- err
		}

		resChan <- res
	}()

	for {
		select {
		case err := <-errChan:
			return nil, err
		case res := <-resChan:
			return res, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
