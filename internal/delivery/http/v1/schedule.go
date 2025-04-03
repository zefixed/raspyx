package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
	"strings"
)

type scheduleRoutes struct {
	uc  *usecase.ScheduleUseCase
	log *slog.Logger
}

// NewScheduleRouteCreate
// @Summary Creating a new schedule
// @Description Creates a new schedule in the database and returns its uuid
// @Tags schedule
// @Accept json
// @Produce json
// @Param schedule body dto.CreateScheduleRequest true "Schedule"
// @Success 200 {object} ResponseOK{response=dto.CreateScheduleRequest}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules [post]
func NewScheduleRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteCreate"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.POST("/", func(c *gin.Context) {
		var scheduleDTO dto.CreateScheduleRequest
		if err := c.ShouldBindJSON(&scheduleDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &scheduleDTO)
		if err != nil {
			errorMapping := []struct {
				contains string
				message  string
				key      string
				value    any
			}{
				{"GroupRepository", "Group not found", "group_number", scheduleDTO.Group},
				{"SubjectTypeRepository", "Subject type not found", "subject_type", scheduleDTO.Type},
				{"SubjectRepository", "Subject not found", "subject_uuid", scheduleDTO.SubjectUUID},
				{"LocationRepository", "Location not found", "location_name", scheduleDTO.Location},
				{"RoomRepository", "Rooms not found", "rooms_number", scheduleDTO.Rooms},
				{"TeacherRepository", "Teachers not found", "teachers_uuid", scheduleDTO.TeachersUUID},
				{"invalid start time", "Invalid start time", "start_time", scheduleDTO.StartTime},
				{"invalid end time", "Invalid end time", "end_time", scheduleDTO.EndTime},
				{"invalid start date", "Invalid start date", "start_date", scheduleDTO.StartDate},
				{"invalid end date", "Invalid end date", "end_date", scheduleDTO.EndDate},
				{"invalid UUID", "Invalid UUID", "invalid_uuid", scheduleDTO},
				{"invalid weekday", "Invalid weekday", "invalid_weekday", scheduleDTO.Weekday},
			}

			for _, em := range errorMapping {
				if strings.Contains(err.Error(), em.contains) {
					log.Info(em.message, slog.Any(em.key, em.value))
					c.JSON(http.StatusBadRequest, RespError(em.message))
					return
				}
			}

			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewScheduleRouteGet
// @Summary Getting schedules
// @Description Get all schedules from database
// @Tags schedule
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules [get]
func NewScheduleRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGet"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/", func(c *gin.Context) {
		resp, err := r.uc.Get(c)
		if err != nil {
			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByUUID
// @Summary Getting schedule by uuid
// @Description Get schedule from database with given uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Schedule uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/uuid/{uuid} [get]
func NewScheduleRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Error("Schedule not found", slog.String("error", err.Error()))
				c.JSON(http.StatusBadRequest, RespError("Schedule not found"))
			} else if strings.Contains(err.Error(), "invalid uuid") {
				log.Error("Invalid uuid", slog.String("invalid_uuid", reqUUID))
				c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}
