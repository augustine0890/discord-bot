package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username"`
	Channel        string             `bson:"channel"`
	Text           string             `bson:"text"`
	Sentiment      string             `bson:"sentiment"`
	SentimentScore map[string]float64 `bson:"sentiment_score"`
	Topic          string             `bson:"topic"`
	CreatedAt      primitive.DateTime `bson:"createdAt"`
	// UpdatedAt      primitive.DateTime `bson:"updatedAt"`
}
