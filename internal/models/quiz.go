package models

import "time"

type Category string

type Quiz struct {
	QuizID         string    `json:"quizId"`
	CreatorUserID  string    `json:"creatorUserId"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Category       Category  `json:"category"`
	ImageURL       string    `json:"imageUrl"`
	DurationSecond int       `json:"durationSecond"`
	CreatedAt      time.Time `json:"createdAt"`
}
