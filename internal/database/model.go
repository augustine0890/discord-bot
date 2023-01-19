package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID                   primitive.ObjectID `bson:"_id"`
	Username             string             `bson:"username"`
	Channel              string             `bson:"channel"`
	Text                 string             `bson:"text"`
	Sentiment            string             `bson:"sentiment"`
	SentimentHuggingFace string             `bson:"huggingface,omitempty"`
	Emotion              string             `bson:"emotion,omitempty"`
	// SentimentScore map[string]float64 `bson:"sentiment_score,omitempty"`
	// Topic          string             `bson:"topic,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
}
