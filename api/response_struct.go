package api

type actor struct {
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	Gravater_id   string `json:"gravatar_id"`
	Url           string `json:"url"`
	Avatar_url    string `json:"avatar_url"`
}

type repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type page struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Action   string `json:"string"`
	Sha      string `json:"sha"`
	Html_url string `json:"html_url"`
}

type payload struct {
	// Event: CommitCommentEvent, DisscussionEvent, ForkEvent, IssueCommentEvent, IssuesEvent, MemberEvent
	Action string `json:"action"`
	// +PullRequestReviewEvent
	Comment map[string]any `json:"comment"`

	// Event: DisscussionEvent
	Discussion map[string]any `json:"discussion"`

	// Event: ForkEvent
	Forkee map[string]any `json:"forkee"`

	// Event: GollumEvent
	Pages []page `json:"pages"`

	// IssueCommentEvent, IssuesEvent
	Issue map[string]any `json:"issue"`

	// IssuesEvent, PullRequestEvent
	Assignee  map[string]any   `json:"assignee"`
	Assignees []map[string]any `json:"assignees"`
	Label     map[string]any   `json:"label"`

	// MemberEvent
	Member map[string]any `json:"member"`

	// PublicEvent
	// payload is empty

	// PullRequestEvent, PullRequestReviewEvent
	Number      int              `json:"number"`
	PullRequest map[string]any   `json:"pull_request"`
	Labels      []map[string]any `json:"labels"`

	// Event: CreateEvent, DeleteEvent
	Ref           string `json:"ref"`
	Ref_type      string `json:"ref_type"`
	Full_ref      string `json:"full_ref"`
	Master_branch string `json:"master_branch"`
	Description   string `json:"description"`
	Pusher_type   string `json:"pusher_type"`

	// PushEvent
	Repo_id int    `json:"repository_id"`
	Push_id int    `json:"push_id"`
	Head    string `json:"head"`
	Before  string `json:"before"`
}

// JSON response struct
type Response struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Actor      actor   `json:"actor"`
	Repo       repo    `json:"repo"`
	Payload    payload `json:"payload"`
	Public     bool    `json:"public"`
	Created_at string  `json:"created_at"`
}
