package handlers

import (
	"net/http"
	"quizy-be/internal/config"
	"quizy-be/internal/models"
	"quizy-be/internal/repository"
	"quizy-be/internal/service"
	"quizy-be/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var authService *service.AuthService

func InitAuthHandlers(db *mongo.Database) {
	emailCfg, err := config.NewEmailConfig()
	if err != nil {
		panic("Email config error: " + err.Error())
	}

	emailSender := func(to []string, subject, htmlBody string) error {
		return utils.SendHTMLEmail(*emailCfg, to, subject, htmlBody)
	}

	signupRepo := repository.NewSignupRequestRepository(db)
	authService = service.NewAuthService(signupRepo, emailSender)
}

func Signup(w http.ResponseWriter, r *http.Request) error {
	var req models.SignupRequest

	if err := utils.ParseJSONBody(r, &req); err != nil {
		return utils.ErrInvalidInput("Invalid request body")
	}

	if err := authService.CreateSignupRequest(r.Context(), &req); err != nil {
		return err
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]string{
		"message": "Verification email sent",
	})
	return nil
}
