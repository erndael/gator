package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("usage: login <username>")
	}
	username := cmd.arguments[0]
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("%s: %v", cmd.name, err)
	}
	fmt.Printf("successfully set %s as the current user", username)
	return nil
}
