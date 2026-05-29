package internal

import (
	"fmt"
	"github-activity/models"
)

func ResponseHandle(responses []models.Response) {
	for _, r := range responses {
		switch r.Type {
		case "CommitCommentEvent":
			fmt.Printf("A commit comment %v  is created to %v by %v", r.Payload.Comment.Body, r.Repo.Name, r.Actor.Display_login)
		case "CreateEvent":
			fmt.Printf("Created %v %v  in %v", r.Payload.RefType, r.Payload.Ref, r.Repo.Name)
		case "DeleteEvent":
			fmt.Printf("Deleted %v %v in %v", r.Payload.RefType, r.Payload.Ref, r.Repo.Name)
		case "ForkEvent":
			fmt.Printf("Forked in %v", r.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("Commented issue %v in %v", r.Payload.Issue.Id, r.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("%v issue %v in %v", r.Payload.Action, r.Payload.Issue.Id, r.Repo.Name)
		case "PushEvent":
			fmt.Printf("Pushed %v to %v", r.Payload.PushId, r.Repo.Name)
		case "PullRequestEvent":

		case "PullRequestReviewEvent":
			//
		case "PullRequestReviewCommentEvent":
			//
		case "WatchEvent":
			//
		default:
			fmt.Printf("%v in %v", r.Type, r.Repo)
		}
	}

	// commitReposCouunt := len(commits)
	// if commitReposCouunt > 0 {
	// 	for repository, commitCount := range commits {
	// 		fmt.Printf("Pushed %v commits to %v\n", commitCount, repository)
	// 	}
	// }

}
