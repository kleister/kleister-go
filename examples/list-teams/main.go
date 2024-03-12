package main

import (
	"log"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-go/kleister"
	"github.com/kleister/kleister-go/kleister/team"
)

func main() {
	transport := client.New(
		"localhost:8080",
		kleister.DefaultBasePath,
		[]string{"http"},
	)

	resp, err := kleister.New(
		transport,
		strfmt.Default,
	).Team.ListTeams(
		team.NewListTeamsParams(),
		client.BasicAuth(
			"admin",
			"admin",
		),
	)

	if err != nil {
		log.Fatalf("Failed to get teams: %s", err)
	}

	for _, t := range resp.Payload.Teams {
		log.Println(kleister.PtrString(t.Name))
	}
}
