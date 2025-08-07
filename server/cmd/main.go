package main

import (
	"log"

	"github.com/Akshat-z/Chat-app/db"
)

func main() {
	_, err := db.New()
	if err != nil {
		log.Fatalf("could not initialize db connection: %s", err)
	}

}
