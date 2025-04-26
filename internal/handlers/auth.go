package handler

import (
	"net/http"
	"quizy-be/pkg/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := utils.ParseJSONBody(r, &loginRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Login successful",
	})
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest SignupRequest
	if err := utils.ParseJSONBody(r, &signupRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
}
