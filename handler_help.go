package main

import (
	"fmt"
)

const helpHelp string = "Prints this help message"

func handlerHelp(s *state, cmd command) error {
	fmt.Println("Available Commands:")
	for key, c := range s.cmds.registeredCommands {
		fmt.Printf("	%-15s %s\n", key, c.help)
	}
	return nil
}
