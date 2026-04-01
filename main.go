package main

import _ "github.com/lib/pq"
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/beerded/gator/internal/config"
	"github.com/beerded/gator/internal/database"
)

func main() {
	fmt.Println("Starting Program")
	defer fmt.Println("Exiting Program")

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments. Need at least 1")
	}

	// cfg.Print()
	db, err := sql.Open("postgres", cfg.DBUrl)

	dbQueries := database.New(db)

	s := state{cfg: cfg, db: dbQueries}

	cmds := newCommands()

 	args := os.Args[1:]
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerCreateFeedFollow))
  	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollowsForUser))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	commandStruct := command{name: args[0], args: args[1:]}

	err = cmds.run(&s, commandStruct)
	if err != nil {
		log.Fatalf("Error running command '%s': %v", args[0], err)
	}
	// cfg.Print()
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error getting user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
