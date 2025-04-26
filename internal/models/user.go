package models

type User struct {
	UID            string  `json:"uid"`
	Email          string  `json:"email"`
	HashedPassword string  `json:"hashedPassword"`
	DisplayName    *string `json:"displayName,omitempty"`
	PhotoURL       string  `json:"photoUrl"`
}
