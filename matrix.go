package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matrix-org/gomatrix"
)

type Config struct {
	MatrixUrl   string `json:"matrixUrl"`
	MatrixRoom  string `json:"matrixRoom"`
	MatrixToken string `json:"matrixToken"`
}

func matrix(pubDate string, Description string, Title string) {
	// Open the config file
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()
	// Decode the JSON file into a Config struct
	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
	// Create a new client with a Matrix server URL and access token
	client, err := gomatrix.NewClient(config.MatrixUrl, config.MatrixRoom, config.MatrixToken)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	// Need to join room to be able to send message, doh
	_, err = client.JoinRoom(config.MatrixRoom, "", nil)
	if err != nil {
		fmt.Println("Error joining room ", err)
		return

	}
	// Send a message to a Matrix room
	var message string
	message = Title + "\n" + Description

	_, err = client.SendText(config.MatrixRoom, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

}
