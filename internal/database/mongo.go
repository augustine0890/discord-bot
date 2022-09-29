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
	UsersCollection      *mongo.Collection
	ActivitiesCollection *mongo.Collection
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

	UsersCollection = client.Database("discord-bot").Collection("users")
	ActivitiesCollection = client.Database("discord-bot").Collection("activities")
}
