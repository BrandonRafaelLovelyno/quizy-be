package models

import "time"

type QuizAttempt struct {
	QuizID             string    `json:"quizId"`
	UserID             string    `json:"userId"`
	CorrectAnswerCount int       `json:"correctAnswerCount"`
	FalseAnswerCount   int       `json:"falseAnswerCount"`
	CreatedAt          time.Time `json:"createdAt"`
}
