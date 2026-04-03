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


	cmds := newCommands()

 	args := os.Args[1:]
	cmds.register("login", handlerLogin, helpLogin)
	cmds.register("register", handlerRegister, helpRegister)
	cmds.register("reset", handlerReset, helpReset)
	cmds.register("users", handlerUsers, helpUsers)
	cmds.register("agg", handlerAgg, helpAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed), helpAddFeed)
	cmds.register("feeds", handlerListFeeds, helpFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerCreateFeedFollow), helpFollow)
  	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollowsForUser), helpFollowing)
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow), helpUnfollow)
	cmds.register("browse", middlewareLoggedIn(handlerBrowse), helpBrowse)
	cmds.register("help", handlerHelp, helpHelp)

	s := state{cfg: cfg, db: dbQueries, cmds: cmds}
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
