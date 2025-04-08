package dto

import (
	"github.com/google/uuid"
)

type ScheduleRequest struct {
	Group        string   `json:"group" example:"221-352"`
	TeachersUUID []string `json:"teachers_uuid,omitempty" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9,b444b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Rooms        []string `json:"rooms,omitempty" example:"ав4805,ав4810"`
	SubjectUUID  string   `json:"subject" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Type         string   `json:"type" example:"Практика"`
	Location     string   `json:"location" example:"Автозаводская"`
	StartTime    string   `json:"start_time" example:"09:00:00"`
	EndTime      string   `json:"end_time" example:"10:30:00"`
	StartDate    string   `json:"start_date" example:"2025-02-01"`
	EndDate      string   `json:"end_date" example:"2025-06-01"`
	Weekday      int      `json:"weekday" example:"1"`
	Link         string   `json:"link" example:"https://rasp.dmami.ru"`
}

type CreateScheduleResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type Week struct {
	Monday    Day `json:"monday"`
	Tuesday   Day `json:"tuesday"`
	Wednesday Day `json:"wednesday"`
	Thursday  Day `json:"thursday"`
	Friday    Day `json:"friday"`
	Saturday  Day `json:"saturday"`
}

type Day struct {
	First   []Pair `json:"1"`
	Second  []Pair `json:"2"`
	Third   []Pair `json:"3"`
	Fourth  []Pair `json:"4"`
	Fifth   []Pair `json:"5"`
	Sixth   []Pair `json:"6"`
	Seventh []Pair `json:"7"`
}

type Pair struct {
	Subject   string   `json:"subject" example:"Иностранный язык"`
	Teachers  []string `json:"teachers" example:"Фамилия Имя Отчество,Фамилия Имя"`
	StartDate string   `json:"start_date" example:"2025-02-01"`
	EndDate   string   `json:"end_date" example:"2025-06-01"`
	Rooms     []string `json:"rooms,omitempty" example:"ав4805,ав4810"`
	Location  string   `json:"location" example:"Автозаводская"`
	Type      string   `json:"type" example:"Практика"`
}
