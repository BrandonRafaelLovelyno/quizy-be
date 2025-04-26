package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupRequest struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	Token        string             `json:"-" bson:"token"`
	ExpiresAt    time.Time          `json:"-" bson:"expiresAt"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
}
