package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beerded/gator/internal/config"
)

func main() {
	fmt.Println("Starting Program")
	defer fmt.Println("Exiting Program")

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments. Need at least 1")
	}

	// cfg.Print()
	s := state{cfg: cfg}

	cmds := newCommands()

 	args := os.Args[1:]
	cmds.register("login", handlerLogin)

	commandStruct := command{name: args[0], args: args[1:]}

	err = cmds.run(&s, commandStruct)
	if err != nil {
		log.Fatalf("Error running command '%s': %v", args[0], err)
	}
	// cfg.Print()
}
