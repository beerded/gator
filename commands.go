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
	cmds	*commands
}

type command struct {
	name    string
    args    []string
}

type commandInfo struct {
	callback		func(*state, command) error
	help			string
}

type commands struct {
	registeredCommands	map[string]commandInfo
}

func newCommands() *commands {
	cmds := commands{registeredCommands: make(map[string]commandInfo)}
	return &cmds
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("Error: empty state")
	}
	if c == nil {
		return errors.New("Commands struct is pointing to nil")
	}
	commando, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("Could not find command '%s'", cmd.name)
	}
	callback := commando.callback
	err := callback(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error, helpText string) {
	//fmt.Printf("Registering command '%s'\n", name)
	c.registeredCommands[name] = commandInfo{
		callback:		f,
		help:			helpText,
	}
}
