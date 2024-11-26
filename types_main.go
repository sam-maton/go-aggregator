package main

import "github.com/sam-maton/go-aggregator/internal/config"

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(state *state, cmd command) error
}

type state struct {
	config *config.Config
}
