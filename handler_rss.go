package main

import (
	"context"
	"fmt"
	"time"

	"github.com/erndael/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user %s: %w\n", s.cfg.CurrentUserName, err)

	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	_, err = s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %w\n", err)
	}

	fmt.Printf("successfully added feed %s\n", name)
	return nil
}
