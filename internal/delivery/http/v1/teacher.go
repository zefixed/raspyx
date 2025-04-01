package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"raspyx/internal/domain/models"
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
// @Tags teacher
// @Accept json
// @Produce json
// @Param teacher body dto.CreateTeacherRequest true "Teacher"
// @Success 200 {object} ResponseOK{response=dto.CreateTeacherRequest}
// @Failure 400 {object} ResponseError
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
			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewTeacherRouteGet
// @Summary Getting teachers
// @Description Get all teachers from database
// @Tags teacher
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetTeachersResponse}
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
			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewTeacherRouteGetByUUID
// @Summary Getting teacher by uuid
// @Description Get teacher from database with given uuid
// @Tags teacher
// @Accept */*
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Success 200 {object} ResponseOK{response=models.Teacher}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/uuid/{uuid} [get]
func NewTeacherRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		teacherUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		resp, err := r.uc.GetByUUID(c, teacherUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Teacher not found", slog.String("teacher_uuid", teacherUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Teacher not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewTeacherRouteGetByFullName
// @Summary Getting teacher by fullname
// @Description Get teacher from database with given fullname
// @Tags teacher
// @Accept */*
// @Produce json
// @Param fullname path string true "Teacher fullname"
// @Success 200 {object} ResponseOK{response=models.Teacher}
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
			if strings.Contains(err.Error(), "not found") {
				log.Info("Teachers not found", slog.String("teacher_fullname", reqFullName))
				c.JSON(http.StatusNotFound, RespError("Teacher not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewTeacherRouteUpdate
// @Summary Updating teacher
// @Description Update teacher in database
// @Tags teacher
// @Accept json
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Param teacher body dto.UpdateTeacherRequest true "Teacher"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/uuid/{uuid} [put]
func NewTeacherRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		teacherUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		var teacherDTO dto.UpdateTeacherRequest
		if err := c.ShouldBindJSON(&teacherDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err = r.uc.Update(c, &models.Teacher{
			UUID:       teacherUUID,
			FirstName:  teacherDTO.FirstName,
			SecondName: teacherDTO.SecondName,
			MiddleName: teacherDTO.MiddleName,
		})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Teacher not found", slog.String("teacher_uuid", teacherUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Teacher not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewTeacherRouteDelete
// @Summary Deleting existing teacher
// @Description Deleting existing teacher from the database
// @Tags teacher
// @Accept */*
// @Produce json
// @Param uuid path string true "Teacher uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/teachers/{uuid} [delete]
func NewTeacherRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.TeacherUseCase, log *slog.Logger) {
	r := &teacherRoutes{uc, log}

	teacherGroup := apiV1Group.Group("/teachers")

	teacherGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		teacherUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, teacherUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Teacher not found", slog.String("teacher_uuid", teacherUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Teacher not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
