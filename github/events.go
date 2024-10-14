package github

type Event struct {
	Id         int         `json:"id"`
	Type       string      `json:"type"`
	Agent      Actor       `json:"actor"`
	Repository Repo        `json:"repository"`
	Load       Payload     `json:"payload"`
	Public     bool        `json:"public"`
	CreatedAt  CreatedTime `json:"created_at"`
}

type Actor struct {
}

type Repo struct {
}

type Payload struct {
}

type Events []Event

type CreatedTime string
