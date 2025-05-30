package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswer struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuestionID primitive.ObjectID `json:"questionId" bson:"questionId"`
	ChoiceID   primitive.ObjectID `json:"choiceId" bson:"choiceId"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
}
