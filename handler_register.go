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

const helpRegister string = "Register a new user and switch to that user"
func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Username is required for the register command")
	}

	username := cmd.args[0]
	uid := uuid.New()
	currentTime := time.Now()

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if user.ID != uuid.Nil {
		// Found an entry in the database. That name must already exist.
  		log.Fatal("User with that name already exists. Exiting")
	}

	user, err = s.db.CreateUser(ctx, database.CreateUserParams{
		ID: 		uid,
		CreatedAt:	currentTime,
		UpdatedAt:	currentTime,
		Name:		username,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Created user %s:\n", username)
	fmt.Printf("	ID: 			%v\n", user.ID)
	fmt.Printf("	CreatedAt: 		%v\n", user.CreatedAt)
	fmt.Printf("	UpdatedAt:  		%v\n", user.UpdatedAt)
	fmt.Printf("	Name: 			%v\n", user.Name)
	fmt.Println("==================================================")
	return nil
}
