package frankerfacez

import (
	"encoding/json"
	"github.com/Back1ng/go-frankerfacez/utils"
	"io"
	"net/http"
	"net/url"
)

type ApiV1EmotesRequest struct {
	Query     string `json:"q,omitempty"`
	Owner     string `json:"owner,omitempty"`
	Artist    string `json:"artist,omitempty"`
	Sensitive bool   `json:"sensitive,omitempty"`
	Sort      string `json:"sort,omitempty"`
	Page      int    `json:"page,omitempty"`
	PerPage   int    `json:"per_page,omitempty"`
}

type ApiV1EmotesResponse struct {
	Pages     int64      `json:"_pages"`
	Total     int64      `json:"_total"`
	Emoticons []Emoticon `json:"emoticons"`
}

func (c Client) GetEmotes(in ApiV1EmotesRequest) (*ApiV1EmotesResponse, error) {
	u, err := url.Parse("https://api.frankerfacez.com/v1/emotes")
	if err != nil {
		return nil, err
	}

	values, err := utils.ReqToQueryValues(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	res, err := c.Get(*u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ApiV1EmotesResponse
	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c Client) GetEmotesByName(name string) (*ApiV1EmotesResponse, error) {
	u, err := url.Parse("https://api.frankerfacez.com/v1/emotes")
	if err != nil {
		return nil, err
	}

	qu := u.Query()
	qu.Set("q", name)
	qu.Set("per_page", "1")
	qu.Set("sort", "count-desc")

	u.RawQuery = qu.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ApiV1EmotesResponse
	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
