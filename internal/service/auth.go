package service

import (
	"context"
	"quizy-be/internal/helper"
	"quizy-be/internal/models"
	"quizy-be/internal/repository"
)

type AuthService struct {
	signupRepo  *repository.SignupRequestRepository
	emailSender func(to []string, subject, htmlBody string) error
}

func NewAuthService(
	signupRepo *repository.SignupRequestRepository,
	emailSender func(to []string, subject, htmlBody string) error,
) *AuthService {
	return &AuthService{
		signupRepo:  signupRepo,
		emailSender: emailSender,
	}
}

func (s *AuthService) CreateSignupRequest(ctx context.Context, req *models.SignupRequest) error {
	if err := helper.EnsureNoDuplicateEmail(ctx, s.signupRepo, req.Email); err != nil {
		return err
	}

	helper.InitSignupRequest(req)

	if err := helper.SendVerificationEmail(s.emailSender, req.Email, req.Token); err != nil {
		return err
	}

	return s.signupRepo.Create(ctx, req)
}
