package main

import (
	"context"
	"fmt"
	"time"

	"github.com/erndael/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("usage: login <username>")
	}
	username := cmd.arguments[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}
	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("%s: %v", cmd.name, err)
	}
	fmt.Printf("successfully set %s as the current user\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("usage: register <username>")
	}
	username := cmd.arguments[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("successfully registered user %s\n", user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully reset users")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		username := user.Name
		if s.cfg.CurrentUserName == username {
			fmt.Printf("* %s (current)\n", username)
			continue
		}
		fmt.Printf("* %s\n", username)
	}
	return nil
}
