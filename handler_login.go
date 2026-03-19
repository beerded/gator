package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required for login command")
	}
	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User set to %s\n", username)

	return nil
}

