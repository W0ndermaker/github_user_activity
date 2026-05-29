package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github-activity/api"
)

func main() {
	client := &http.Client{}

	url := "https://api.github.com/users/nilbuild/events"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp := make([]api.Response, 0)

	if err = json.Unmarshal(body, &resp); err != nil {
		fmt.Println(err)
	}

	for _, v := range resp {
		printResponse(v)
	}
}

func printResponse(r api.Response) {
	// fmt.Printf("\nId : %v\n", r.Id)
	fmt.Printf("Type : %v\n", r.Type)
	// fmt.Printf("Actor : %v\n", r.Actor)
	// fmt.Printf("Repo : %v\n", r.Repo)
	// fmt.Printf("Payload: \n")
	// printPayload(r)
	fmt.Printf("Payload: %v\n", r.Payload)
	// fmt.Printf("Public : %v\n", r.Public)
	// fmt.Printf("Created_at : %v\n", r.Created_at)

}

func printPayload(r api.Response) {
	fmt.Printf("\tAction: %v\n", r.Payload.Action)
	fmt.Printf("\tRef: %v\n", r.Payload.Ref)
	fmt.Printf("\tPush_id: %v\n", r.Payload.Push_id)
	fmt.Printf("\tBefore: %v\n", r.Payload.Before)
	fmt.Printf("\tRepo_id: %v\n", r.Payload.Repo_id)
}
