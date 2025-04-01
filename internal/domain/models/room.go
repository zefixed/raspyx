package models

import "github.com/google/uuid"

type Room struct {
	UUID   uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Number string    `json:"number" example:"ав4805"`
}
