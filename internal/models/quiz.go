package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category string

type Quiz struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	CreatorID   primitive.ObjectID   `json:"creatorId" bson:"creatorId"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Category    Category             `json:"category" bson:"category"`
	ImageURL    string               `json:"imageUrl" bson:"imageUrl"`
	Duration    int                  `json:"duration" bson:"duration"`
	Questions   []primitive.ObjectID `json:"questions" bson:"questions"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt" bson:"updatedAt"`
}
