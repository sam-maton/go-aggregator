package main

import (
	"context"
	"fmt"

	"github.com/sam-maton/go-aggregator/internal/database"
)

func middlewareLogin(handler func(state *state, cmd command, user database.User) error) func(*state, command) error {
	return func(state *state, cmd command) error {
		user, err := state.db.GetUser(context.Background(), state.config.UserName)

		if err != nil {
			return fmt.Errorf("there was an error getting the user: %w", err)
		}

		return handler(state, cmd, user)
	}
}
