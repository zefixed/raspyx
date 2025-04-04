package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
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
// @Param schedule body dto.ScheduleRequest true "Schedule"
// @Success 200 {object} ResponseOK{response=dto.ScheduleRequest}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules [post]
func NewScheduleRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteCreate"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.POST("/", func(c *gin.Context) {
		var scheduleDTO dto.ScheduleRequest
		if err := c.ShouldBindJSON(&scheduleDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &scheduleDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "schedule_dto",
				logValue: scheduleDTO,
			})
			return
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
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
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
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "error",
				logValue: err.Error(),
			})
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
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
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
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "schedule_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByTeacher
// @Summary Getting schedule by teacher fullname
// @Description Get schedule from database with given teacher fullname
// @Tags schedule
// @Accept */*
// @Produce json
// @Param fn path string true "Teacher fullname"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/teacher/fn/{fn} [get]
func NewScheduleRouteGetByTeacher(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByTeacher"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/teacher/fn/:fn", func(c *gin.Context) {
		reqfn := c.Param("fn")
		resp, err := r.uc.GetByTeacher(c, reqfn)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "teacher_fullname",
				logValue: reqfn,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByTeacherUUID
// @Summary Getting schedule by teacher uuid
// @Description Get schedule from database with given teacher uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/teacher/uuid/{uuid} [get]
func NewScheduleRouteGetByTeacherUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByTeacherUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/teacher/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByTeacherUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "teacher_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByGroup
// @Summary Getting schedule by group number
// @Description Get schedule from database with given group number
// @Tags schedule
// @Accept */*
// @Produce json
// @Param number path string true "Group number"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/group/number/{number} [get]
func NewScheduleRouteGetByGroup(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByGroup"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/group/number/:number", func(c *gin.Context) {
		reqNumber := c.Param("number")
		resp, err := r.uc.GetByGroup(c, reqNumber)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "group_number",
				logValue: reqNumber,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByGroupUUID
// @Summary Getting schedule by group uuid
// @Description Get schedule from database with given group uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Group uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/group/uuid/{uuid} [get]
func NewScheduleRouteGetByGroupUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByGroupUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedules")

	scheduleGroup.GET("/group/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByGroupUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "group_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByRoom
// @Summary Getting schedule by room number
// @Description Get schedule from database with given room number
// @Tags schedule
// @Accept */*
// @Produce json
// @Param number path string true "Room number"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/room/number/{number} [get]
func NewScheduleRouteGetByRoom(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByRoom"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleRoom := apiV1Group.Group("/schedules")

	scheduleRoom.GET("/room/number/:number", func(c *gin.Context) {
		reqNumber := c.Param("number")
		resp, err := r.uc.GetByRoom(c, reqNumber)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "room_number",
				logValue: reqNumber,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByRoomUUID
// @Summary Getting schedule by room uuid
// @Description Get schedule from database with given room uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Room uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/room/uuid/{uuid} [get]
func NewScheduleRouteGetByRoomUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByRoomUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleRoom := apiV1Group.Group("/schedules")

	scheduleRoom.GET("/room/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByRoomUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "room_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetBySubject
// @Summary Getting schedule by subject name
// @Description Get schedule from database with given subject name
// @Tags schedule
// @Accept */*
// @Produce json
// @Param name path string true "Subject name"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/subject/name/{name} [get]
func NewScheduleRouteGetBySubject(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetBySubject"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleSubject := apiV1Group.Group("/schedules")

	scheduleSubject.GET("/subject/name/:name", func(c *gin.Context) {
		reqName := c.Param("name")
		resp, err := r.uc.GetBySubject(c, reqName)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_name",
				logValue: reqName,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetBySubjectUUID
// @Summary Getting schedule by subject uuid
// @Description Get schedule from database with given subject uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Subject uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/subject/uuid/{uuid} [get]
func NewScheduleRouteGetBySubjectUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetBySubjectUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleSubject := apiV1Group.Group("/schedules")

	scheduleSubject.GET("/subject/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetBySubjectUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByLocation
// @Summary Getting schedule by location name
// @Description Get schedule from database with given location name
// @Tags schedule
// @Accept */*
// @Produce json
// @Param name path string true "Location name"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/location/name/{name} [get]
func NewScheduleRouteGetByLocation(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByLocation"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleLocation := apiV1Group.Group("/schedules")

	scheduleLocation.GET("/location/name/:name", func(c *gin.Context) {
		reqName := c.Param("name")
		resp, err := r.uc.GetByLocation(c, reqName)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "location_name",
				logValue: reqName,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteGetByLocationUUID
// @Summary Getting schedule by location uuid
// @Description Get schedule from database with given location uuid
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Location uuid"
// @Success 200 {object} ResponseOK{response=dto.Week}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedules/location/uuid/{uuid} [get]
func NewScheduleRouteGetByLocationUUID(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewScheduleRouteGetByLocationUUID"
	log = log.With(slog.String("op", op))

	r := &scheduleRoutes{uc, log}

	scheduleLocation := apiV1Group.Group("/schedules")

	scheduleLocation.GET("/location/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByLocationUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "location_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewScheduleRouteUpdate
// @Summary Updating room
// @Description Update room in database
// @Tags schedule
// @Accept json
// @Produce json
// @Param uuid path string true "Schedule uuid"
// @Param room body dto.ScheduleRequest true "Schedule"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedule/uuid/{uuid} [put]
func NewScheduleRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedule")

	scheduleGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var scheduleDTO dto.ScheduleRequest
		if err := c.ShouldBindJSON(&scheduleDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &scheduleDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "schedule",
				logValue: map[string]any{"uuid": reqUUID, "schedule_dto": scheduleDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewScheduleRouteDelete
// @Summary Deleting existing schedule
// @Description Deleting existing schedule from the database
// @Tags schedule
// @Accept */*
// @Produce json
// @Param uuid path string true "Schedule uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/schedule/{uuid} [delete]
func NewScheduleRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.ScheduleUseCase, log *slog.Logger) {
	r := &scheduleRoutes{uc, log}

	scheduleGroup := apiV1Group.Group("/schedule")

	scheduleGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "schedule_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
