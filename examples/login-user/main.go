package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-go/v1/kleister"
	"github.com/kleister/kleister-go/v1/kleister/auth"
	"github.com/kleister/kleister-go/v1/models"
)

func main() {
	username := "your-username"
	password := strfmt.Password("your-password")

	resp, err := kleister.Default.Auth.LoginUser(
		auth.NewLoginUserParams().WithAuthLogin(
			&models.AuthLogin{
				Username: &username,
				Password: &password,
			},
		),
	)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
