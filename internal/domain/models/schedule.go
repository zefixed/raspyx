package models

import (
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	UUID         uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	GroupUUID    uuid.UUID `json:"group_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	SubjectUUID  uuid.UUID `json:"subject_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	TypeUUID     uuid.UUID `json:"type_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	LocationUUID uuid.UUID `json:"location_uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	StartTime    time.Time `json:"start_time" example:"09:00:00"`
	EndTime      time.Time `json:"end_time" example:"10:30:00"`
	StartDate    time.Time `json:"start_date" example:"2025-02-01"`
	EndDate      time.Time `json:"end_date" example:"2025-06-10"`
	Weekday      int       `json:"weekday" example:"1"`
	Link         string    `json:"link" example:"https://rasp.dmami.ru"`
	IsSession    bool      `json:"isSession" example:"false"`
}

type ScheduleData struct {
	UUID      uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Group     string    `json:"group" example:"221-352"`
	Teachers  []string  `json:"teachers,omitempty" example:"Фамилия Имя Отчество,Фамилия Имя"`
	Rooms     []string  `json:"rooms,omitempty" example:"ав4805,ав4810"`
	Subject   string    `json:"subject" example:"Иностранный язык"`
	Type      string    `json:"type" example:"Практика"`
	Location  string    `json:"location" example:"Автозаводская"`
	StartTime time.Time `json:"start_time" example:"09:00:00"`
	EndTime   time.Time `json:"end_time" example:"10:30:00"`
	StartDate time.Time `json:"start_date" example:"2025-02-01"`
	EndDate   time.Time `json:"end_date" example:"2025-06-01"`
	Weekday   int       `json:"weekday" example:"1"`
	Link      string    `json:"link" example:"https://rasp.dmami.ru"`
}
