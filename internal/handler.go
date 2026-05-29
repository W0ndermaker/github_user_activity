package internal

import (
	"github-activity/api"
)

func responseHandle(responses []api.Response) {
	for _, r := range responses {
		switch r.Type {

		case "CommitCommentEvent":
			commitCommentEvent()
		case "CreateEvent":
			createEvent(r)
		case "DeleteEvent":
		case "DisscussionEvent":
		case "ForkEvent":
		case "GollumEvent":
		case "IssueCommentEvent":
		case "IssuesEvent":
		case "MemberEvent":
		case "PublicEvent":
		case "PushEvent":
		case "PullRequestEvent":
		case "PullRequestReviewEvent":
		case "PullRequestReviewCommentEvent":
		case "ReleaseEvent":
		case "WatchEvent":
			watchEvent()
		}
	}
}

func createEvent(r api.Response) {

}

func watchEvent() {
}

func commitCommentEvent() {}
