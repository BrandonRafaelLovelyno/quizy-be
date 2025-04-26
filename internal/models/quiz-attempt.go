package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizAttempt struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID   `json:"userId" bson:"userId"`
	QuizID     primitive.ObjectID   `json:"quizId" bson:"quizId"`
	Score      float64              `json:"score" bson:"score"`
	Answers    []primitive.ObjectID `json:"answers" bson:"answers"`
	StartedAt  time.Time            `json:"startedAt" bson:"startedAt"`
	FinishedAt *time.Time           `json:"finishedAt,omitempty" bson:"finishedAt,omitempty"`
}
