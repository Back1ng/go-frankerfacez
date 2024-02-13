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
	c := frankerfacez.NewClient(ctx)

	res, err := c.GetEmotes(frankerfacez.ApiV1EmotesRequest{
		Query:    "KEKW",
		Animated: true,
		PerPage:  1,
		Sort:     frankerfacez.SORT_COUNT_DESC,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Emoticons[0].Animated)
}
