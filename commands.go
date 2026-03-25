package main

import (
	"errors"
	"fmt"

    "github.com/beerded/gator/internal/config"
    "github.com/beerded/gator/internal/database"
)

type state struct{
	cfg     *config.Config
	db		*database.Queries
}

type command struct {
	name    string
    args    []string
}

type commands struct {
	registeredCommands	map[string]func(*state, command) error
}

func newCommands() *commands {
	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	return &cmds
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("Error: empty state")
	}
	if c == nil {
		return errors.New("Commands struct is pointing to nil")
	}
	callback, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("Could not find command '%s'", cmd.name)
	}
	err := callback(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	//fmt.Printf("Registering command '%s'\n", name)
	c.registeredCommands[name] = f
}
