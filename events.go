package main

import (
	"fmt"
	"io"
	"net/http"
)

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

	fmt.Println("body", string(body))

	return nil
}
