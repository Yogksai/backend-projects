package main

import (
	"project/activity"
)

func main() {
	// Example usage
	username := "bradfitz"
	activities, err := activity.GetGithubUserActivity(username)
	if err != nil {
		panic(err)
	}
	activity.DisplayActivity(username, activities)
}
