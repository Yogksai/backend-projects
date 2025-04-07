package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/josn"
	model "project/models"

	"github.com/charmbracelet/lipgloss"
)

func GetGithubUserActivity(username string) ([]model.GithubUserActivity, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, fmt.Errorf("user not found, %s", response.Status)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user activity, %s", response.Status)
	}
	var activities []model.GithubUserActivity
	if err := json.NewDecoder(response.Body).Decode(&activities); err != nil {
		return nil, err
	}
	josn.WriteToJson(activities)
	return activities, nil

}

func DisplayActivity(username string, events []model.GithubUserActivity) error {
	if len(events) == 0 {
		return fmt.Errorf("no activity found")
	}

	for _, event := range events {
		var action string
		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commits to %s", commitCount, event.Repo.Name)

		case "IssuesEvent":
			action = fmt.Sprintf("%s issue in %s", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Create %s in %s", event.Payload.RefType, event.Repo.Name)
		default:
			action = fmt.Sprintf("%s inn %s", event.Type, event.Repo.Name)
		}

		actionStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("- %s", action))
		fmt.Println(actionStyle)
	}

	return nil
}
