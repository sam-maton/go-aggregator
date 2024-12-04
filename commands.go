package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sam-maton/go-aggregator/internal/database"
	"github.com/sam-maton/go-aggregator/internal/rss"
)

func loginHandler(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the login command requires at least one argument")
	}

	username := cmd.args[0]

	user, err := state.db.GetUser(context.Background(), username)

	if err != nil {
		fmt.Println("There was an issue logging in the user:")
		return err
	}

	err = state.config.SetUser(user.Name)

	if err != nil {
		return err
	}

	fmt.Println("Welecome " + username + "! You were logged in successfully.")
	return nil
}

func registerUserHandler(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the register command requires at least one argument")
	}

	username := cmd.args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := state.db.CreateUser(context.Background(), params)

	if err != nil {
		fmt.Println("There was an issue registering the user:")
		return err
	}

	err = state.config.SetUser(user.Name)

	if err != nil {
		fmt.Println("There was an issue updating the config after registering the user:")
		return err
	}

	fmt.Println("The new user " + user.Name + " was successfully created.")

	return nil
}

func usersHandler(state *state, cmd command) error {
	users, err := state.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("there was an error whilst trying to get the list of users: %w", err)
	}

	for _, v := range users {
		if v.Name == state.config.UserName {
			fmt.Printf("* %s (current)\n", v.Name)
		} else {
			fmt.Printf("* %s\n", v.Name)
		}
	}
	return nil
}

func resetHandler(state *state, cmd command) error {
	err := state.db.DeleteUsers(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Print("Database successfully reset")
	return nil
}

func aggHandler(state *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("there was an error whilst getting the RSS feed: %w", err)
	}

	fmt.Println(feed)

	return nil
}

func addFeedHandler(state *state, cmd command, user database.User) error {

	if len(cmd.args) < 2 {
		return errors.New("the addFeed command requires at least two arguments")
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := state.db.CreateFeed(context.Background(), params)

	if err != nil {
		return fmt.Errorf("there was an error when creating the new feed: %w", err)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	_, err = state.db.CreateFeedFollow(context.Background(), followParams)

	if err != nil {
		return fmt.Errorf("there was an error when following the new feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}

func listFeedsHandler(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("there was an error getting the list of Feeds: %w", err)
	}

	for _, f := range feeds {
		fmt.Println(f.Name, f.Url, f.UserName.String)
	}

	return nil
}

func followHandler(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("the follow command requires at least one argument")
	}

	url := cmd.args[0]

	feed, err := state.db.GetFeedByUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("there was an error getting the feed: %w", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollowRow, err := state.db.CreateFeedFollow(context.Background(), args)

	if err != nil {
		return fmt.Errorf("there was an error whilst following the feed: %w", err)
	}

	fmt.Println(feedFollowRow.UserName + " successfully followed " + feedFollowRow.FeedName)

	return nil
}

func followingHandler(state *state, cmd command, user database.User) error {

	feeds, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("there was an error getting the users feeds: %w", err)
	}

	if len(feeds) < 1 {
		fmt.Println("You are not following any feeds! Use the 'follow' command to follow a feed.")

		return nil
	}
	fmt.Println("You are currently following these feeds:")
	for _, f := range feeds {
		fmt.Println("* " + f.FeedName)
	}

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
