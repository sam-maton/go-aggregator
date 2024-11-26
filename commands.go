package main

import (
	"errors"
	"fmt"
)

func loginHandler(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the login command requires at least one argument")
	}

	username := cmd.args[0]
	err := state.config.SetUser(username)

	if err != nil {
		return err
	}

	fmt.Println("Username " + username + " was set.")
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, exists := c.commandMap[cmd.name]

	if !exists {
		return errors.New("command '" + cmd.name + "' does not exist")
	}

	err := cmdFunc(s, cmd)

	if err != nil {
		return err
	}

	return nil
}
