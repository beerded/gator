package main

import (
	"fmt"
	"log"
	"context"

	"github.com/google/uuid"
)

const helpLogin string = "Login as a different user (must have been previously added)"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required for login command")
	}
	username := cmd.args[0]

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if user.ID == uuid.Nil {
		log.Fatalf("User '%s' not found in database. Unable to login. Exiting\n", username)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User set to %s\n", username)

	return nil
}

