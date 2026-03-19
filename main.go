package main

import (
	"fmt"
	"github.com/beerded/gator/internal/config"
)

func main() {
	fmt.Println("Starting Program")
	defer fmt.Println("Exiting Program")

	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("%w", err)
	}
	cfg.Print()
	err = cfg.SetUser("Mr. Foobar")
	if err != nil {
		fmt.Errorf("%w", err)
	}
	cfg.Print()
}
