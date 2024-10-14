package github

type Event struct {
	Id         int         `json:"id"`
	Type       string      `json:"type"`
	Agent      Actor       `json:"actor"`
	Repository Repo        `json:"repository"`
	Public     bool        `json:"public"`
	CreatedAt  CreatedTime `json:"created_at"`
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

type CreatedTime string
