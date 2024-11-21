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

	fmt.Println(c.DatabaseURL)

}
