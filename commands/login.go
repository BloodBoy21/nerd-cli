package commands

import (
	"fmt"
	"golang.org/x/term"
	"nerd-cli/helpers"
	"os"
)

type CredentialResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func Login(command *CommandService) {
	config, error := helpers.LoadConfig()
	if error != nil {
		fmt.Println("Error loading config:", error)
		return
	}
	username, password := requireCredentials()
	loginRequest := helpers.Request{
		Url:     "login",
		Payload: helpers.ParsePayload(map[string]interface{}{"email": username, "password": password}),
	}
	data := loginRequest.CallAPI("POST")
	var response CredentialResponse
	helpers.ParseJson(data, &response)
	if response.AccessToken == "" {
		fmt.Println("Login failed")
		fmt.Println("Error:", data)
		return
	}
	config.TOKEN = response.AccessToken
	error = helpers.SaveConfigValue("token", response.AccessToken)
	if error != nil {
		panic(error)
	}
	fmt.Println("Login successful")
}

func requireCredentials() (string, string) {
	var username string
	var password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		return "", ""
	}
	password = string(bytePassword)

	fmt.Println()

	return username, password
}
