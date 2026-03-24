package main

import (
	"context"
	"errors"
	"fmt"
  	"log"
	"time"

	"github.com/google/uuid"
	"github.com/beerded/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Username is required for the register command")
	}

	username := cmd.args[0]
	uid := uuid.New()
	currentTime := time.Now()
	args := database.CreateUserParams{
		ID: 		uid,
		CreatedAt:	currentTime,
		UpdatedAt:	currentTime,
		Name:		username,
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if user.ID != uuid.Nil {
		// Found an entry in the database. That name must already exist.
  		log.Fatal("User with that name already exists. Exiting")
	}

	user, err = s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Created user %s:\n  ID: %v\n  CreatedAt:  %v\n  UpdatedAt:  %v\n  Name: %v\n", username, user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
