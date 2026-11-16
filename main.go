package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/erndael/gator/internal/config"
	"github.com/erndael/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
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

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	registry := commands{
		commands: map[string]func(*state, command) error{},
	}

	registry.register("login", handlerLogin)
	registry.register("register", handlerRegister)
	registry.register("reset", handlerReset)
	registry.register("users", handlerUsers)
	registry.register("agg", handlerAgg)

	if err := registry.run(s, cmd); err != nil {
		log.Fatal(err)
	}

}
