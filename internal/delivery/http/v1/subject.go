package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"raspyx/internal/domain/models"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type subjectRoutes struct {
	uc  *usecase.SubjectUseCase
	log *slog.Logger
}

// NewSubjectRouteCreate
// @Summary Creating a new subject
// @Description Creates a new subject in the database and returns its uuid
// @Tags subject
// @Accept json
// @Produce json
// @Param subject body dto.CreateSubjectRequest true "Subject name"
// @Success 200 {object} ResponseOK{response=dto.CreateSubjectRequest}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects [post]
func NewSubjectRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewSubjectRouteCreate"
	log = log.With(slog.String("op", op))

	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.POST("/", func(c *gin.Context) {
		var subjectDTO dto.CreateSubjectRequest
		if err := c.ShouldBindJSON(&subjectDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &subjectDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_dto",
				logValue: subjectDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewSubjectRouteGet
// @Summary Getting subjects
// @Description Get all subjects from database
// @Tags subject
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetSubjectsResponse}
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects [get]
func NewSubjectRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewSubjectRouteGet"
	log = log.With(slog.String("op", op))

	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.GET("/", func(c *gin.Context) {
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

// NewSubjectRouteGetByUUID
// @Summary Getting subject by uuid
// @Description Get subject from database with given uuid
// @Tags subject
// @Accept */*
// @Produce json
// @Param uuid path string true "Subject uuid"
// @Success 200 {object} ResponseOK{response=models.Subject}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects/uuid/{uuid} [get]
func NewSubjectRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		subjectUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		resp, err := r.uc.GetByUUID(c, subjectUUID)
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

// NewSubjectRouteGetByName
// @Summary Getting subject by name
// @Description Get subject from database with given name
// @Tags subject
// @Accept */*
// @Produce json
// @Param name path string true "Subject name"
// @Success 200 {object} ResponseOK{response=[]models.Subject}
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects/name/{name} [get]
func NewSubjectRouteGetByName(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.GET("/name/:name", func(c *gin.Context) {
		reqName := c.Param("name")

		resp, err := r.uc.GetByName(c, reqName)
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

// NewSubjectRouteUpdate
// @Summary Updating subject
// @Description Update subject in database
// @Tags subject
// @Accept json
// @Produce json
// @Param uuid path string true "Subject uuid"
// @Param subject body dto.UpdateSubjectRequest true "Subject"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects/uuid/{uuid} [put]
func NewSubjectRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		subjectUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		var subjectDTO dto.UpdateSubjectRequest
		if err := c.ShouldBindJSON(&subjectDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err = r.uc.Update(c, &models.Subject{UUID: subjectUUID, Name: subjectDTO.Name})
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject",
				logValue: map[string]any{"uuid": reqUUID, "subject_dto": subjectDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewSubjectRouteDelete
// @Summary Deleting existing subject
// @Description Deleting existing subject from the database
// @Tags subject
// @Accept */*
// @Produce json
// @Param uuid path string true "Subject uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjects/{uuid} [delete]
func NewSubjectRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.SubjectUseCase, log *slog.Logger) {
	r := &subjectRoutes{uc, log}

	subjectGroup := apiV1Group.Group("/subjects")

	subjectGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		subjectUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, subjectUUID)
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

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
