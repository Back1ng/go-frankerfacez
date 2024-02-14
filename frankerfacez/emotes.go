package frankerfacez

import (
	"context"
	"encoding/json"
	"github.com/Back1ng/go-frankerfacez/utils"
	"io"
	"net/http"
	"net/url"
)

// Sorting constants using in requests
const (
	SORT_NAME_ASC     = "name-asc"
	SORT_NAME_DESC    = "name-desc"
	SORT_OWNER_ASC    = "owner-asc"
	SORT_OWNER_DESC   = "owner-desc"
	SORT_COUNT_ASC    = "count-asc"
	SORT_COUNT_DESC   = "count-desc"
	SORT_UPDATED_ASC  = "updated-asc"
	SORT_UPDATED_DESC = "updated-desc"
	SORT_CREATED_ASC  = "created-asc"
	SORT_CREATED_DESC = "created-desc"
)

type ApiV1EmotesRequest struct {
	// Query A string to search by
	Query string `json:"q"`

	// Owner A string to search users by
	Owner string `json:"owner,omitempty"`

	// Artist A string to search users by
	Artist string `json:"artist,omitempty"`

	// Sensitive Whether the search query should be treated as case-sensitive.
	Sensitive bool `json:"sensitive,omitempty"`

	// Sort The column and direction to sort by.
	// Possible values: name-asd, name-desc, owner-asc, owner-desc,
	// count-asc, count-desc, updated-asc, updated-desc, created-asc, created-desc
	// Can be used from const SORT_...
	Sort string `json:"sort,omitempty"`

	Animated bool `json:"animated,omitempty"`

	Page int `json:"page,omitempty"`

	// PerPage Number of emotes per page. Range: 1-200
	PerPage int `json:"per_page,omitempty"`
}

type ApiV1EmotesResponse struct {
	// Pages total count of pages
	Pages int64 `json:"_pages"`

	// Total count of emotes
	Total int64 `json:"_total"`

	// Emoticons founded emotes
	Emoticons []Emoticon `json:"emoticons"`
}

// GetEmotes implements /v1/emotes handler, that get emotes in bulk
func (c *client) GetEmotes(ctx context.Context, in ApiV1EmotesRequest) (*ApiV1EmotesResponse, error) {
	u, err := url.Parse("https://api.frankerfacez.com/v1/emotes")
	if err != nil {
		return nil, err
	}

	values, err := utils.ReqToQueryValues(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	res, err := c.Get(ctx, *u)
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

// GetEmotesByName implements /v1/emotes handler, that select emotes by name
// order by usage_count and get one emote
func (c *client) GetEmotesByName(name string) (*ApiV1EmotesResponse, error) {
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
