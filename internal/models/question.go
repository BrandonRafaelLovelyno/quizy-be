package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuizID    primitive.ObjectID `json:"quizId" bson:"quizId"`
	Number    int                `json:"number" bson:"number"`
	ImageURL  *string            `json:"imageUrl,omitempty"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
