package stashClient

type Pr struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       string     `json:"state"`
	Open        bool       `json:"open"`
	Closed      bool       `json:"closed"`
	FromRef     Ref        `json:"fromRef"`
	ToRef       Ref        `json:"ToRef"`
	Locked      bool       `json:"locked"`
	Reviewers   []Reviewer `json:"reviewers"`
}

type Ref struct {
	ID         string `json:"id"`
	Repository Repo   `json:"repository"`
}

type Repo struct {
	Slug    string      `json:"slug"`
	Name    interface{} `json:"name"`
	Project Proj        `json:"project"`
}

type Proj struct {
	Key string `json:"key"`
}

type Reviewer struct {
	User Usr `json:"user"`
}

type Usr struct {
	Name string `json:"name"`
}
