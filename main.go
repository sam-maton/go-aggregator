package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	run()
}

func run() {
	mainState := setupState()
	mainCommands := setupCommands()

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments were provided.")

		os.Exit(1)
	}

	splitCommand := command{
		name: args[1],
		args: args[2:],
	}

	err := mainCommands.run(&mainState, splitCommand)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
