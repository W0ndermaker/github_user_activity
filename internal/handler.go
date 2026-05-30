package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserAcitivity(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error in http request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("github API error: status code %v", response.StatusCode)
	}

	events := make([]GithubActivity, 0)
	if err := json.NewDecoder(response.Body).Decode(&events); err != nil {
		return fmt.Errorf("json decoding error: %v", err)
	}

	eventsHandle(events)

	return nil
}

func eventsHandle(githubActivities []GithubActivity) {
	if len(githubActivities) == 0 {
		fmt.Println("No recent activities")
		return
	}
	for _, r := range githubActivities {
		switch r.Type {
		case "CommitCommentEvent":
			fmt.Printf("A commit comment %v  is created to %v by %v\n", r.Payload.Comment.Body, r.Repo.Name, r.Actor.Display_login)
		case "CreateEvent":
			fmt.Printf("Created %v %v  in %v\n", r.Payload.RefType, r.Payload.Ref, r.Repo.Name)
		case "DeleteEvent":
			fmt.Printf("Deleted %v %v in %v\n", r.Payload.RefType, r.Payload.Ref, r.Repo.Name)
		case "ForkEvent":
			fmt.Printf("Forked in %v\n", r.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("Commented issue %v in %v\n", r.Payload.Issue.Id, r.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("%v issue %v in %v\n", r.Payload.Action, r.Payload.Issue.Id, r.Repo.Name)
		case "PushEvent":
			fmt.Printf("Pushed %v to %v\n", r.Payload.PushId, r.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("%v %v on %v\n", r.Payload.Action, r.Payload.PullReq.Url, r.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %v\n", r.Repo.Name)
		default:
			fmt.Printf("%v in %v\n", r.Type, r.Repo.Name)
		}
	}
}
