package main

import (
	"fmt"
	"os"

	"github.com/sam-maton/go-aggregator/internal/config"
)

func main() {
	c, err := config.Read()

	if err != nil {
		fmt.Println(err)
	}

	mainState := state{
		config: &c,
	}

	mainCommands := commands{
		commandMap: map[string]func(state *state, cmd command) error{},
	}

	mainCommands.register("login", loginHandler)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments were provided.")

		os.Exit(1)
	}

	splitCommand := command{
		name: args[1],
		args: args[2:],
	}

	err = mainCommands.run(&mainState, splitCommand)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
