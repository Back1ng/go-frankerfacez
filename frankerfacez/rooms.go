package frankerfacez

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type ApiV1RoomByNameRequest struct {
	Name string `json:"name"`
}

type ApiV1RoomByNameResponse struct {
	Room Room `json:"room"`
	Sets struct {
		Set Set `json:"-"`
	} `json:"sets"`
}

type Room struct {
	ID          int    `json:"_id"`
	TwitchID    int    `json:"twitch_id"`
	IDName      string `json:"id"`
	IsGroup     bool   `json:"is_group"`
	DisplayName string `json:"display_name"`
	Set         int    `json:"set"`
}

type Set struct {
	ID        int        `json:"id"`
	Type      int        `json:"_type"`
	Title     string     `json:"title"`
	Emoticons []Emoticon `json:"emoticons"`
}

func (c *client) GetRoomByName(ctx context.Context, in ApiV1RoomByNameRequest) (*ApiV1RoomByNameResponse, error) {
	u, err := url.Parse(fmt.Sprintf("https://api.frankerfacez.com/v1/room/%v", in.Name))
	if err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, *u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ApiV1RoomByNameResponse
	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	var unknownNameSet map[string]interface{}
	if err := json.Unmarshal(r, &unknownNameSet); err != nil {
		return nil, err
	}

	// set response.Room.Set as key in sets...
	sets := unknownNameSet["sets"].(map[string]interface{})[fmt.Sprint(response.Room.Set)]
	s, _ := json.Marshal(sets)

	if err := json.Unmarshal(s, &response.Sets.Set); err != nil {
		return nil, err
	}

	return &response, nil
}
