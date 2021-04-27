package main

import (
	"fmt"

	"github.com/icedream/go-footballdata"
)

func main() {
	client := new(footballdata.Client).
		WithToken("b3ca26ec78a64948b4086d3f074390d0")

	competitions, err := client.Competitions().Do()
	if err != nil {
		fmt.Println(err)
	}

	for _, competition := range competitions {
		fmt.Println(competition.Id, competition.Caption)
	}

}
