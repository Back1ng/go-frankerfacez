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

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
