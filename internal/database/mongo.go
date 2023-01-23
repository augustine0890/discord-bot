package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// UsersCollection = client.Database("discord-bot").Collection("users")
	// ActivitiesCollection = client.Database("discord-bot").Collection("activities")
	MessagesCollection = client.Database("discord-stats").Collection("messages")
}

func CreateMessage(msg Message, ctx context.Context) (err error) {
	_, insertErr := MessagesCollection.InsertOne(ctx, msg)
	if insertErr != nil {
		log.Println(insertErr)
		return insertErr
	}

	// log.Println("New Message ID:", result.InsertedID)
	return nil
}

func DeleteMessageWeekly() (count int64, deleteErr error) {
	// Only save latest one month data
	oneMonthBefore := time.Now().AddDate(0, -1, 0)
	time := primitive.NewDateTimeFromTime(oneMonthBefore)

	filter := bson.D{{"createdAt", bson.D{{"$lte", time}}}}

	results, deleteErr := MessagesCollection.DeleteMany(context.TODO(), filter)
	if deleteErr != nil {
		log.Println(deleteErr)
		return count, deleteErr
	}
	return results.DeletedCount, nil
}
