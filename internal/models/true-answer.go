package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrueAnswer struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuestionID primitive.ObjectID `json:"questionId" bson:"questionId"`
	ChoiceID   primitive.ObjectID `json:"choiceId" bson:"choiceId"`
}
