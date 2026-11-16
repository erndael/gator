package main

import (
	"context"
	"fmt"
	"html"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(html.UnescapeString(feed.Channel.Title))
	return nil
}
