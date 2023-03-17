package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"

	"discordbot/internal/commands"
	"discordbot/internal/config"
	"discordbot/internal/monitor"
	"discordbot/internal/utils"

	"discordbot/internal/database"
	"discordbot/internal/discord"
)

var ctx = context.TODO()

func main() {
	// Flag will be stored in the stage variable at runtime
	stage := flag.String("stage", "prod", "The enviroment running")
	flag.Parse()

	// Loading enviroment variables
	err := utils.LoadEnv(*stage)
	if err != nil {
		fmt.Printf("Error loading environment: %v\n", err)
	}
	log.Printf("Running with %v enviroment\n", *stage)

	// Reading config file
	// cgf, err := config.ReadConfig()
	// if err != nil {
	// fmt.Println("error reading config file", err)
	// }

	// Connect to Database
	database.Start(ctx)

	// Delete messages weekly
	c := cron.New()
	// Running At 10:00, only on Monday (0 10 * * MON)
	// Every minute (* * * * *)
	c.AddFunc("CRON_TZ=Asia/Seoul 0 10 * * MON", func() {
		count, err := database.DeleteMessageWeekly()
		if err != nil {
			log.Printf("Error deleting messages%v\n", err)
		}
		log.Printf("Deleting %v messages \n", count)
	})
	c.Start()

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Println("Error creating Discord session: ", err)
	}

	// // Register the messageCreate func as a callback for MessageCreate events.
	// dg.AddHandler(messageCreate)

	// registerCommands(dg, cgf)
	dg.AddHandler(discord.SendMessageHandler)
	// Run the monitoring function in a separatre goroutine
	go monitor.StartMonitoring(dg)

	dg.AddHandler(discord.ProcessMessageHandler)
	// Only care about receiving message events
	// dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listenning
	err = dg.Open()
	if err != nil {
		log.Println("Error opening websocket connection: ", err)
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
