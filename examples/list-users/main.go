package main

import (
	"context"
	"log"

	"github.com/kleister/kleister-go/kleister"
)

func main() {
	client, err := kleister.NewClientWithResponses(
		"http://localhost:8080/api/v1",
	)

	if err != nil {
		log.Fatalf("Failed to initialize client: %s", err)
	}

	resp, err := client.ListUsersWithResponse(
		context.Background(),
		&kleister.ListUsersParams{},
	)

	if err != nil {
		log.Fatalf("Failed to get users: %s", err)
	}

	for _, t := range kleister.FromPtr(resp.JSON200.Users) {
		log.Println(kleister.FromPtr(t.Username))
	}
}
