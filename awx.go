package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AzureADOAuth2Settings struct {
	Key     string                 `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_KEY"`
	Secret  string                 `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET"`
	OrgMap  map[string]OrgMapEntry `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_ORGANIZATION_MAP"`
	TeamMap interface{}            `json:"SOCIAL_AUTH_AZUREAD_OAUTH2_TEAM_MAP"`
}

type OrgMapEntry struct {
	Users  bool `json:"users"`
	Admins bool `json:"admins"`
}

func awx(appId string, secret string) {
	url := "http://127.0.0.1:10445/api/v2/settings/azuread-oauth2/"
	username := ""
	password := ""
	authHeader := "Basic " + base64Encode(username+password)

	// Create the request body
	requestBody := AzureADOAuth2Settings{
		Key:    appId,
		Secret: secret,
		OrgMap: map[string]OrgMapEntry{
			"Default": {
				Users:  true,
				Admins: true,
			},
		},
		TeamMap: nil,
	}

	// Marshal the request body into JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP PATCH request
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request was successful.")
	} else {
		fmt.Printf("Request failed with status: %d\n", resp.StatusCode)
	}
}

func base64Encode(s string) string {
	return s // You should implement a proper base64 encoding function here.
	// for the poc I used the curl from postman and let it create that for us
}
