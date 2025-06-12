package main

import (
	"log"
)

func main() {

	username, err := cmdFlags()

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = getUserEvents(username)

	if err != nil {
		log.Fatalf("%v", err)
	}
}
