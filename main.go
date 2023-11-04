package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type client struct {
	AppID        string `json:"appId"`
	TenantID     string `json:"tenant"`
	ClientSecret string `json:"password"`
}

func main() {
	var3 := os.Getenv("client")
	fmt.Println("printing json")
	fmt.Println(var3)

	var cred client

	err := json.Unmarshal([]byte(var3), &cred)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(cred)
	fmt.Println("client secret recieved")

	awx(cred.AppID, cred.ClientSecret)
}
