package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sam-maton/go-aggregator/internal/rss"
)

func main() {
	// run()
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(feed.Channel.Item[0].Description)
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
