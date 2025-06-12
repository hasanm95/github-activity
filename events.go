package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubEvent struct {
	Type    string `json:"type"`
	Repo    struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct{} `json:"commits,omitempty"`
	} `json:"payload"`
}

func getUserEvents(username string) error {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body")
	}

	var events []GitHubEvent
	if err := json.Unmarshal(body, &events); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commits to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		}
	}

	return nil
}
