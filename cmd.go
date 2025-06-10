package main

import (
	"flag"
	"fmt"
)

func cmdFlags() (string, error) {
	// Parse flags
	flag.Parse()

	// Get remaining arguments after flags
	args := flag.Args()

	// Check if any username was provided
	if len(args) == 0 {
		return "", fmt.Errorf("no username provided")
	}

	username := args[0]

	return username, nil
}
