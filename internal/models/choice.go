package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Choice struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuestionID primitive.ObjectID `json:"questionId" bson:"questionId"`
	Letter     rune               `json:"letter" bson:"letter"`
	Body       string             `json:"body" bson:"body"`
}
