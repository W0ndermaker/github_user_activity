package main

import (
	"fmt"
	"log"
	"os"

	"github-activity/internal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one argument: the GitHub username.")
		os.Exit(1)
	}
	username := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/users/%v/events", username)

	if err := internal.GetUserAcitivity(url); err != nil {
		log.Fatal(err)
	}
}
