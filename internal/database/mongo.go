package database

import (
	"context"
	"discordbot/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MessagesCollection *mongo.Collection
	// ActivitiesCollection *mongo.Collection
)

func Start(ctx context.Context) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(config.Database).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// UsersCollection = client.Database("discord-bot").Collection("users")
	// ActivitiesCollection = client.Database("discord-bot").Collection("activities")
	MessagesCollection = client.Database("discord-bot").Collection("messages")
}

func CreateMessage(msg Message, ctx context.Context) (err error) {
	result, insertErr := MessagesCollection.InsertOne(ctx, msg)
	if insertErr != nil {
		log.Println(insertErr)
		return insertErr
	}

	log.Println("New Message ID:", result.InsertedID)
	return nil
}
