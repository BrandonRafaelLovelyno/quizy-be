package middleware

import (
	"errors"
	"net/http"
	"quizy-be/pkg/utils"
)

func ErrorHandler(next func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		var uError *utils.Error
		if errors.As(err, &uError) {
			switch uError.Code {
			case "INVALID_INPUT":
				http.Error(w, uError.Message, http.StatusBadRequest)
				return
			default:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
