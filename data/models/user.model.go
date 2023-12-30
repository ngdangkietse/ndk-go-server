package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
}
