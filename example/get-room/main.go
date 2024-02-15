package main

import (
	"context"
	"fmt"
	"github.com/Back1ng/go-frankerfacez/frankerfacez"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c := frankerfacez.NewClient()

	res, err := c.GetRoomByName(ctx, frankerfacez.ApiV1RoomByNameRequest{
		Name: "silvername",
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, emote := range res.Sets.Set.Emoticons {
		fmt.Println(emote.Name)
	}
}
