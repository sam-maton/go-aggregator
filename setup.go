package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sam-maton/go-aggregator/internal/config"
	"github.com/sam-maton/go-aggregator/internal/database"
)

func setupCommands() commands {
	cmds := commands{
		commandMap: map[string]func(state *state, cmd command) error{},
	}

	cmds.register("login", loginHandler)
	cmds.register("register", registerUserHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", aggHandler)
	cmds.register("addfeed", middlewareLogin(addFeedHandler))
	cmds.register("feeds", listFeedsHandler)
	cmds.register("follow", middlewareLogin(followHandler))
	cmds.register("following", middlewareLogin(followingHandler))
	cmds.register("unfollow", middlewareLogin(unfollowHandler))
	cmds.register("posts", middlewareLogin(postsHandler))

	return cmds
}

func setupState() state {
	c, err := config.Read()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Database setup
	db, err := sql.Open("postgres", c.DatabaseURL)
	dbQueries := database.New(db)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return state{
		db:     dbQueries,
		config: &c,
	}
}
