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

	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't delete feeds: %w", err)
	}
	fmt.Println("Successfully reset the database")
	return nil
}
