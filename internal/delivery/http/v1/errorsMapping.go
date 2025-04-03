package v1

import (
	"fmt"
	"strings"
)

func mapError(err error) string {
	other := []struct {
		e       string
		message string
	}{
		{"RoomsToScheduleRepository", "RoomsToSchedule"},
		{"TeachersToScheduleRepository", "Duplicate teachers"},
		{"invalid UUID", "Invalid UUID"},
		{"invalid start time", "Invalid start time"},
		{"invalid end time", "Invalid end time"},
		{"invalid start date", "Invalid start date"},
		{"invalid end date", "Invalid end date"},
		{"invalid weekday", "Invalid weekday"},
		{"fk error", "Object with given uuid does not exist"},
	}

	for _, o := range other {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(o.e)) {
			return o.message
		}
	}

	repos := []struct {
		repo    string
		message string
	}{
		{"GroupRepository", "Group"},
		{"LocationRepository", "Location"},
		{"RoomRepository", "Room"},
		{"ScheduleRepository", "Schedule"},
		{"SubjectRepository", "Subject"},
		{"SubjectTypeRepository", "Subject type"},
		{"TeacherRepository", "Teacher"},
	}

	for _, r := range repos {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(r.repo)) {
			if strings.Contains(strings.ToLower(err.Error()), "not found") {
				return fmt.Sprintf("%v not found", r.message)
			} else if strings.Contains(strings.ToLower(err.Error()), "exist") {
				return fmt.Sprintf("%v exists", r.message)
			} else if strings.Contains(strings.ToLower(err.Error()), "not exist") {
				return fmt.Sprintf("%v does not exist", r.message)
			}
		}
	}

	return "Unknown error"
}
