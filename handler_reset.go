package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Successfully reset the users table")
	return nil
}
