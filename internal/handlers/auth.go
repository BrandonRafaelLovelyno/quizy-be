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

func Signup(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest

	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := authService.CreateSignupRequest(r.Context(), &req); err != nil {
		switch err {
		case service.ErrEmailExists:
			utils.WriteErrorResponse(w, http.StatusConflict, err.Error())
		default:
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to process signup request")
		}
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]string{
		"message": "Verification email sent",
	})
}
