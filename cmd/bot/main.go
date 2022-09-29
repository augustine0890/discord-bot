package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"discordbot/internal/config"
	"discordbot/internal/database"
)

var ctx = context.TODO()

func main() {
	// Reading config file
	err := config.ReadConfig()

	// Connect to Database
	database.Start(ctx)

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session: ", err)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// Only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listenning
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening websocket connection: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Only care about messages that are "ping"
	// if m.Content != "!ping" {
	// return
	// }

	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "ping!")
	}
}
