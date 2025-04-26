package service

import (
	"context"
	"quizy-be/internal/models"
	"quizy-be/internal/repository"
	"quizy-be/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	signupRepo *repository.SignupRequestRepository
}

func NewAuthService(signupRepo *repository.SignupRequestRepository) *AuthService {
	return &AuthService{signupRepo: signupRepo}
}

func (s *AuthService) CreateSignupRequest(ctx context.Context, req *models.SignupRequest) error {
	existing, err := s.signupRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if existing != nil {
		return ErrEmailExists
	}

	req.ID = primitive.NewObjectID()
	req.Token = utils.GenerateRandomString(32)
	req.ExpiresAt = time.Now().Add(24 * time.Hour)
	req.CreatedAt = time.Now()

	return s.signupRepo.Create(ctx, req)
}

var ErrEmailExists = utils.NewError("verification already sent to this email")
