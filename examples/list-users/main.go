package main

import (
	"log"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-go/kleister"
	"github.com/kleister/kleister-go/kleister/user"
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
	).User.ListUsers(
		user.NewListUsersParams(),
		client.BasicAuth(
			"admin",
			"admin",
		),
	)

	if err != nil {
		log.Fatalf("Failed to get users: %s", err)
	}

	for _, u := range resp.Payload.Users {
		log.Println(kleister.PtrString(u.Username))
	}
}
