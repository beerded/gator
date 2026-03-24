package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't delete users: %w", err)
	}
	fmt.Println("Successfully reset the database")
	return nil
}
