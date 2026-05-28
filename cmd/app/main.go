package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// type actor struct {
// 	Id            int    `json:"id"`
// 	Login         string `json:"login"`
// 	Display_login string `json:"display_login"`
// 	Gravater_id   string `json:"gravatar_id"`
// 	Url           string `json:"url"`
// 	Avatar_url    string `json:"avatar_url"`
// }

type repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// type payload struct {
// 	Repo_id int    `json:"repository_id"`
// 	Push_id int    `json:"push_id"`
// 	Ref     string `json:"ref"`
// 	Head    string `json:"head"`
// 	Before  string `json:"before"`
// }

type response struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	//Actor      actor   `json:"actor"`
	Repo repo `json:"repo"`
	// Payload    payload `json:"payload"`
	// Public     bool    `json:"public"`
	//Created_at string `json:"created_at"`
}

func main() {
	client := &http.Client{}

	url := "https://api.github.com/users/justskiv/events"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(body))
	// fmt.Println("/////////////////////////////////////////////////////")

	var r []response

	if err = json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
