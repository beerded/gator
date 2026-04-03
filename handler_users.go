package main

import (
	"context"
	"fmt"
)

const helpUsers string = "List the available users"

func handlerUsers(s *state, cmd command) error {

	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error getting users: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	fmt.Println("USERS:")
	if len(users) == 0 {
		fmt.Println("Database empty. No users to show")
		return nil
	}

	for _, user := range users {
		toPrint := fmt.Sprintf("* %s", user)
		if user == currentUser {
			toPrint += fmt.Sprintf(" (current)")
		}
		toPrint += "\n"
		fmt.Printf(toPrint)
	}

	return nil
}
