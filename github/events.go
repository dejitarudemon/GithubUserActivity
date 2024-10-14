package github

type Event struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Agent      Actor  `json:"actor"`
	Repository Repo   `json:"repository"`
	Public     bool   `json:"public"`
	CreatedAt  string `json:"created_at"`
}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Events []Event

// Event types
const (
	COMMIT                   = "CommitCommentEvent"
	CREATE                   = "CreateEvent"
	DELETE                   = "DeleteEvent"
	FORK                     = "ForkEvent"
	GOLLUM                   = "GollumEvent"
	ISSUECOMMENT             = "IssueCommentEvent"
	ISSUE                    = "IssueEvent"
	MEMBER                   = "MemberEvent"
	PUBLIC                   = "PublicEvent"
	PULLREQUEST              = "PullRequestEvent"
	PULLREQUESTREVIEW        = "PullRequestReviewEvent"
	PULLREQUESTREVIEWCOMMENT = "PullRequestReviewCommentEvent"
	PULLREQUESTREVIEWTHREAD  = "PullRequestReviewThreadEvent"
	PUSH                     = "PushEvent"
	RELEASE                  = "ReleaseEvent"
	SPONSORSHIP              = "SponsorhipEvent"
	WATCH                    = "WatchEvent"
)
