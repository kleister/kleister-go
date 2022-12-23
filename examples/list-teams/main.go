package main

import (
	"context"
	"log"

	connect "github.com/bufbuild/connect-go"
	"github.com/kleister/kleister-go/kleister"
	teamsv1 "github.com/kleister/kleister-go/kleister/teams/v1"
)

func main() {
	client := kleister.New(
		kleister.WithBaseURL("http://localhost:8080"),
	)

	resp, err := client.Teams.List(
		context.Background(),
		&connect.Request[teamsv1.ListRequest]{},
	)

	if err != nil {
		log.Fatalf("Failed to get teams: %s", err)
	}

	for _, team := range resp.Msg.Teams {
		log.Println(team.Name)
	}
}
