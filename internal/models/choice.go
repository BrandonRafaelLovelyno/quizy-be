package models

type Choice struct {
	QuestionID string `json:"questionId"`
	Letter     rune   `json:"letter"`
	Body       string `json:"body"`
}
