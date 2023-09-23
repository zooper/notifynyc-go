package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func telegram(pubDate string, Title string, Description string) {
	botToken := os.Getenv("botToken")
	channelID := os.Getenv("channgelID")

	// Initialize the bot with your token.
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	var my_message string

	my_message = Title + "\n" + Description
	// Create a message configuration.
	message := tgbotapi.NewMessageToChannel(channelID, my_message)

	// Send the message.
	_, err = bot.Send(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	} else {
		fmt.Println("Message sent successfully to the channel!")
	}
}
