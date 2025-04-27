package utils

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func ErrInvalidInput(msg string) *Error {
	return &Error{
		Code:    "INVALID_INPUT",
		Message: msg,
	}
}

func ErrInternal(msg string) *Error {
	return &Error{
		Code:    "INTERNAL_ERROR",
		Message: msg,
	}
}
