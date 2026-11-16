package main

import (
	"context"
	"database/sql"
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

	user, err := s.db.GetUser(
		context.Background(),
		sql.NullString{
			String: username,
			Valid:  true,
		},
	)
	if err != nil {
		return err
	}
	if err := s.cfg.SetUser(user.Name.String); err != nil {
		return fmt.Errorf("%s: %v", cmd.name, err)
	}
	fmt.Printf("successfully set %s as the current user\n", user.Name.String)
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
		Name: sql.NullString{
			String: username,
			Valid:  true,
		},
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name.String)
	fmt.Printf("successfully register user %s\n", user.Name.String)
	return nil
}
