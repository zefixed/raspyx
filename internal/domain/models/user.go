package models

import "github.com/google/uuid"

type User struct {
	UUID         uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Username     string    `json:"username" example:"username"`
	PasswordHash string    `json:"password_hash"`
	AccessLevel  int       `json:"access_level" example:"0"`
}
