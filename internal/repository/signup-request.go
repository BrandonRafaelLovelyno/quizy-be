package repository

import (
	"context"
	"quizy-be/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignupRequestRepository struct {
	collection *mongo.Collection
}

func NewSignupRequestRepository(db *mongo.Database) *SignupRequestRepository {
	return &SignupRequestRepository{
		collection: db.Collection("signupRequests"),
	}
}

func (r *SignupRequestRepository) FindByEmail(ctx context.Context, email string) (*models.SignupRequest, error) {
	var request models.SignupRequest
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&request)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &request, nil
}

func (r *SignupRequestRepository) Create(ctx context.Context, request *models.SignupRequest) error {
	request.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, request)
	return err
}
