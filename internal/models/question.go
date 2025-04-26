package models

import "time"

type Question struct {
	QuizID    string    `json:"quizId"`
	Number    int       `json:"number"`
	ImageURL  *string   `json:"imageUrl,omitempty"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
