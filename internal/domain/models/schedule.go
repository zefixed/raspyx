package models

import (
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	UUID         uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	TeacherUUID  uuid.UUID `json:"teacher_uuid,omitempty"`
	GroupUUID    uuid.UUID `json:"group_uuid"`
	RoomUUID     uuid.UUID `json:"room_uuid,omitempty"`
	SubjectUUID  uuid.UUID `json:"subject_uuid"`
	TypeUUID     uuid.UUID `json:"type_uuid"`
	LocationUUID uuid.UUID `json:"location_uuid"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Weekday      int       `json:"weekday"`
	Link         string    `json:"link"`
}

type ScheduleData struct {
	UUID      uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Teacher   string    `json:"teacher,omitempty"`
	Group     string    `json:"group"`
	Room      string    `json:"room,omitempty"`
	Subject   string    `json:"subject"`
	Type      string    `json:"type"`
	Location  string    `json:"location"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Weekday   int       `json:"weekday"`
	Link      string    `json:"link"`
}
