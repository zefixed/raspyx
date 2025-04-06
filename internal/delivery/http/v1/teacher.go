package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
	"strings"
)

type teacherRoutes struct {
	uc  *usecase.TeacherUseCase
	log *slog.Logger
}

// NewTeacherRouteCreate
// @Summary Creating a new teacher
// @Description Creates a new teacher in the database and returns its uuid
// @Security ApiKeyAuth
// @Tags teacher
// @Accept json
// @Produce json
// @Param teacher body dto.CreateTeacherRequest true "Teacher"
// @Success 200 {object} ResponseOK{response=dto.CreateTeacherRequest}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers [post]
func NewTeacherRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewTeacherRouteCreate"
	log = log.With(slog.String("op", op))

	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.POST("/", func(c *gin.Context) {
		var teacherDTO dto.CreateTeacherRequest
		if err := c.ShouldBindJSON(&teacherDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &teacherDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "teacher_dto",
				logValue: teacherDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewTeacherRouteGet
// @Summary Getting teachers
// @Description Get all teachers from database
// @Security ApiKeyAuth
// @Tags teacher
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetTeachersResponse}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers [get]
func NewTeacherRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewTeacherRouteGet"
	log = log.With(slog.String("op", op))

	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.GET("/", func(c *gin.Context) {
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

// NewTeacherRouteGetByUUID
// @Summary Getting teacher by uuid
// @Description Get teacher from database with given uuid
// @Security ApiKeyAuth
// @Tags teacher
// @Accept */*
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Success 200 {object} ResponseOK{response=models.Teacher}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/uuid/{uuid} [get]
func NewTeacherRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
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

// NewTeacherRouteGetByFullName
// @Summary Getting teacher by fullname
// @Description Get teacher from database with given fullname
// @Security ApiKeyAuth
// @Tags teacher
// @Accept */*
// @Produce json
// @Param fullname path string true "Teacher fullname"
// @Success 200 {object} ResponseOK{response=models.Teacher}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/fullname/{fullname} [get]
func NewTeacherRouteGetByFullName(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.GET("/fullname/:fullname", func(c *gin.Context) {
		reqFullName := strings.TrimSpace(c.Param("fullname"))

		resp, err := r.uc.GetByFullName(c, reqFullName)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "teacher_fullname",
				logValue: reqFullName,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewTeacherRouteUpdate
// @Summary Updating teacher
// @Description Update teacher in database
// @Security ApiKeyAuth
// @Tags teacher
// @Accept json
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Param teacher body dto.UpdateTeacherRequest true "Teacher"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/{uuid} [put]
func NewTeacherRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.PUT("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var teacherDTO dto.UpdateTeacherRequest
		if err := c.ShouldBindJSON(&teacherDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &teacherDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "teacher",
				logValue: map[string]any{"uuid": reqUUID, "teacher_dto": teacherDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewTeacherRouteDelete
// @Summary Deleting existing teacher
// @Description Deleting existing teacher from the database
// @Security ApiKeyAuth
// @Tags teacher
// @Accept */*
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/{uuid} [delete]
func NewTeacherRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
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

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
