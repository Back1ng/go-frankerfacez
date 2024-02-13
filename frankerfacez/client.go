package frankerfacez

import (
	"context"
	"net/http"
	"net/url"
)

type Client interface {
	Get(url url.URL) (*http.Response, error)

	ApiV1
}

type ApiV1 interface {
	GetEmotes(in ApiV1EmotesRequest) (*ApiV1EmotesResponse, error)
}

type client struct {
	ctx context.Context

	client http.Client
}

func NewClient(ctx context.Context) Client {
	return &client{
		ctx:    ctx,
		client: http.Client{},
	}
}

func (c *client) Get(url url.URL) (*http.Response, error) {
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
		case <-c.ctx.Done():
			return nil, c.ctx.Err()
		}
	}
}
