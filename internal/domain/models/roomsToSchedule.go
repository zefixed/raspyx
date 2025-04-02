package models

import "github.com/google/uuid"

type RoomsToSchedule struct {
	RoomUUID     uuid.UUID `json:"room_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	ScheduleUUID uuid.UUID `json:"schedule_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}
