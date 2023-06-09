package main

import (
	"fmt"
	"os"

	"github.com/matrix-org/gomatrix"
)

func matrix(pubDate string, Description string, Title string) {

	MatrixUrl := os.Getenv("matrixUrl")
	MatrixRoom := os.Getenv("matrixRoom")
	MatrixToken := os.Getenv("matrixToken")

	// Create a new client with a Matrix server URL and access token
	client, err := gomatrix.NewClient(MatrixUrl, MatrixRoom, MatrixToken)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	// Need to join room to be able to send message, doh
	_, err = client.JoinRoom(MatrixRoom, "", nil)
	if err != nil {
		fmt.Println("Error joining room ", err)
		return

	}
	// Send a message to a Matrix room
	var message string
	message = Title + "\n" + Description

	_, err = client.SendText(MatrixRoom, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

}
