package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"discordbot/internal/commands"
	"discordbot/internal/config"

	"discordbot/internal/database"
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
