package main

import (
	"fmt"
	"log"
)

func main() {

	username, err := cmdFlags()

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Main func", username)
}
