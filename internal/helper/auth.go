package helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"quizy-be/internal/models"
	"quizy-be/internal/templates"
	"quizy-be/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.SignupRequest, error)
}

type EmailSender func([]string, string, string) error

func EnsureNoDuplicateEmail(ctx context.Context, repo SignupRepository, email string) error {
	existing, err := repo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existing != nil {
		return ErrEmailExists
	}

	return nil
}

func InitSignupRequest(req *models.SignupRequest) {
	req.ID = primitive.NewObjectID()
	req.Token = utils.GenerateRandomString(32)
	req.ExpiresAt = time.Now().Add(24 * time.Hour)
	req.CreatedAt = time.Now()
}

func SendVerificationEmail(emailSender EmailSender, email, token string) error {
	verificationLink := fmt.Sprintf("https://yourapp.com/verify?token=%s", token)
	emailBody := fmt.Sprintf(templates.SignupEmailTemplate, verificationLink)

	if err := emailSender([]string{email}, "Verify Your Email", emailBody); err != nil {
		log.Printf("Failed to send verification email: %v", err)
		return err
	}
	return nil
}

var ErrEmailExists = utils.NewError("verification already sent to this email")
