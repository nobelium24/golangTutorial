package github

import "time"

const IssueUrl = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json "total_count"`
	Items      []*Issue
}

type Issue struct {
	User      *User
	HTMLURL   string `json: "html_url"`
	Title     string
	State     string
	Number    int
	CreatedAt time.Time `json: "created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json: "html_url"`
}
