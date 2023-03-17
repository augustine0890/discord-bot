package discord

import (
	"context"
	"discordbot/internal/database"
	"discordbot/internal/sentiment"
	"discordbot/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shirou/gopsutil/v3/mem"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

func ProcessMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("The error: %v", r)
			log.Println(err.Error())
		}
	}()
	// Get the channel information
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// If getting the channel information fails, fall back to using API (this may cause rate limiting)
		channel, err = s.Channel(m.ChannelID)
		if err != nil {
			fmt.Printf("Error getting channel info: %v\n", err)
			return
		}
	}
	// Ignore direct messages (DMs)
	if channel.Type == discordgo.ChannelTypeDM {
		return
	}

	// Check users
	if utils.IgnoreUser(m.Author.ID) {
		return
	}

	// Get channel
	if utils.IgnoreChannel(channel.ID) {
		return
	}

	content, err := m.ContentWithMoreMentionsReplaced(s)
	if err != nil {
		log.Printf("Error getting mentions replaced content %s", err.Error())
		content = m.Content
	}

	// Check valid message content
	err = utils.IsValidContent(content)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// fmt.Println("Sever's id", m.GuildID, 1019782712799805440) //

	awsClient := sentiment.NewAwsClient()
	result, err := awsClient.DetectSentiment(content)
	if err != nil {
		log.Println("Detect sentiment error: ", err)
		return
	}
	// Get KST
	kst := m.Timestamp.Add(time.Hour * 9)

	contentReq := &sentiment.TextRequest{Text: content}

	// Sentiment analysis with huggingface
	huggingFaceRes, err := sentiment.HuggingFaceSentiment(*contentReq, "sentiment")
	// log.Println("LABEL", huggingFaceRes.Label)

	// Emotion classification
	// emotion, err := sentiment.HuggingFaceSentiment(*contentReq, "emotion")

	if err != nil {
		msg := database.Message{
			ID:        primitive.NewObjectID(),
			Username:  m.Author.Username,
			Channel:   channel.Name,
			Text:      content,
			Sentiment: *result.Sentiment,
			CreatedAt: primitive.NewDateTimeFromTime(kst),
		}

		err = database.CreateMessage(msg, ctx)
		if err != nil {
			log.Println("Create MongoDB Messsage: ", err)
			return
		}

		return
	}

	// Sentiment Score
	// var ss map[string]float64
	// data, _ := json.Marshal(result.SentimentScore)
	// json.Unmarshal(data, &ss)

	msg := database.Message{
		ID:                   primitive.NewObjectID(),
		Username:             m.Author.Username,
		Channel:              channel.Name,
		Text:                 content,
		Sentiment:            *result.Sentiment,
		SentimentHuggingFace: huggingFaceRes.Label,
		// Emotion:              emotion.Label,
		CreatedAt: primitive.NewDateTimeFromTime(kst),
	}

	err = database.CreateMessage(msg, ctx)
	if err != nil {
		return
	}
}

func SendMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func SendAlertEmbedMessage(s *discordgo.Session, v *mem.VirtualMemoryStat) error {
	embed := &discordgo.MessageEmbed{
		Title: "Discord BOT Alert",
		Color: 0xff0000, // Red color.
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Total memory",
				Value:  fmt.Sprintf("%.2f GB", float64(v.Total)/(1024*1024*1024)),
				Inline: true,
			},
			{
				Name:   "Used memory",
				Value:  fmt.Sprintf("%.2f GB", float64(v.Used)/(1024*1024*1024)),
				Inline: true,
			},
			{
				Name:   "Free memory",
				Value:  fmt.Sprintf("%.2f GB", float64(v.Free)/(1024*1024*1024)),
				Inline: true,
			},
			{
				Name:   "Used memory percentage",
				Value:  fmt.Sprintf("%.2f%%", v.UsedPercent),
				Inline: true,
			},
		},
	}

	// userID := "623155071735037982"
	userID := "1026733912778625026"
	// Send a DM to the new member with the long text
	dm, err := s.UserChannelCreate(userID)
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(dm.ID, embed)
	if err != nil {
		return err
	}
	return nil
}
