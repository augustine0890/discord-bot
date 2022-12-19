package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"discordbot/internal/commands"
	"discordbot/internal/config"
	"discordbot/internal/utils"

	"discordbot/internal/database"
	"discordbot/internal/sentiment"
)

var ctx = context.TODO()

func main() {
	// Reading config file
	cgf, err := config.ReadConfig()
	if err != nil {
		fmt.Println("error reading config file", err)
	}
	// Connect to Database
	database.Start(ctx)

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + cgf.Token)
	if err != nil {
		fmt.Println("error creating Discord session: ", err)
	}

	// // Register the messageCreate func as a callback for MessageCreate events.
	// dg.AddHandler(messageCreate)

	registerCommands(dg, cgf)

	dg.AddHandler(messageHandler)
	// Only care about receiving message events
	// dg.Identify.Intents = discordgo.IntentsGuildMessages

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

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	s.AddHandler(cmdHandler.HandleMessage)
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	err := utils.IsValidContent(content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Sever's id", m.GuildID, 1019782712799805440) //
	awsClient := sentiment.NewAwsClient()
	result, err := awsClient.DetectSentiment(content)
	if err != nil {
		fmt.Println("Detect sentiment error: ", err)
		return
	}

	// Get channel
	channel, _ := s.Channel(m.ChannelID)

	// Sentiment Score
	var ss map[string]float64
	data, _ := json.Marshal(result.SentimentScore)
	json.Unmarshal(data, &ss)

	msg := database.Message{
		ID:             primitive.NewObjectID(),
		Username:       m.Author.Username,
		Channel:        channel.Name,
		Text:           content,
		Sentiment:      *result.Sentiment,
		SentimentScore: ss,
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now()),
	}

	err = database.CreateMessage(msg, ctx)
	if err == nil {
		return
	}
}
