package main

import (
	"fmt"

	"github.com/sam-maton/go-aggregator/internal/config"
)

func main() {
	c, err := config.Read()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.UserName)

	c.SetUser("test_user")

	fmt.Println(c.UserName)

}
