package internal

type actor struct {
	Display_login string `json:"display_login"`
}

type repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type comment struct {
	Body string `json:"body"`
}

type issue struct {
	Id int `json:"id"`
}

type pullRequest struct {
	Url string `json:"url"`
}

type payload struct {
	Action  string      `json:"action"`
	Comment comment     `json:"comment"`
	RepoId  int         `json:"repository_id"`
	PushId  int         `json:"push_id"`
	Issue   issue       `json:"issue"`
	Ref     string      `json:"ref"`
	RefType string      `json:"ref_type"`
	PullReq pullRequest `json:"pull_request"`
}

type GithubActivity struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Actor   actor   `json:"actor"`
	Repo    repo    `json:"repo"`
	Payload payload `json:"payload"`
}
