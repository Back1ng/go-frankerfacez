package frankerfacez

import (
	"context"
	"net/http"
	"net/url"
)

type Client struct {
	ctx context.Context

	client http.Client
}

func NewClient(ctx context.Context) *Client {
	return &Client{
		ctx:    ctx,
		client: http.Client{},
	}
}

func (c *Client) Get(url url.URL) (*http.Response, error) {
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
