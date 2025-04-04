package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

type ErrResp struct {
	err      error
	c        *gin.Context
	log      *slog.Logger
	logKey   string
	logValue any
}

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
		{"invalid fullname", "Invalid fullname"},
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
		{"SubjectRepository", "Subject"},
		{"SubjectTypeRepository", "Subject type"},
		{"TeacherRepository", "Teacher"},
		{"ScheduleRepository", "Schedule"},
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

func makeErrResponse(er *ErrResp) {
	errMes := mapError(er.err)
	if errMes != "Unknown error" {
		er.log.Info(errMes, slog.Any(er.logKey, er.logValue))
		if strings.Contains(errMes, "not found") {
			er.c.JSON(http.StatusNotFound, RespError(errMes))
		} else {
			er.c.JSON(http.StatusBadRequest, RespError(errMes))
		}
		return
	}
	er.log.Error("Internal server error", slog.String("error", er.err.Error()))
	er.c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
	return
}
