package main

import (
	"log"
	"os"

	"github.com/erndael/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("usage: gator <command>")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}

	s := &state{
		cfg: &cfg,
	}

	registry := commands{
		commands: map[string]func(*state, command) error{},
	}
	registry.register("login", handlerLogin)
	if err := registry.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
