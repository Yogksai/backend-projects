package model

type GithubUserActivity struct {
	Type string
	Repo struct {
		Name string
	}
	Payload struct {
		Action  string
		Ref     string // branch name
		RefType string
		Commits []struct {
			Message string
		}
	}
	Created_at string
}
