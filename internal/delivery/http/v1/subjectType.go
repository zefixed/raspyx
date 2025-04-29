package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type subjectTypeRoutes struct {
	uc  *usecase.SubjectTypeUseCase
	log *slog.Logger
}

// NewSubjectTypeRouteCreate
// @Summary Creating a new subjectType
// @Description Creates a new subjectType in the database and returns its uuid
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept json
// @Produce json
// @Param subjectType body dto.CreateSubjectTypeRequest true "Subject type"
// @Success 200 {object} ResponseOK{response=dto.CreateSubjectTypeRequest}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes [post]
func NewSubjectTypeRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewSubjectTypeRouteCreate"
	log = log.With(slog.String("op", op))

	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.POST("/", func(c *gin.Context) {
		var subjectTypeDTO dto.CreateSubjectTypeRequest
		if err := c.ShouldBindJSON(&subjectTypeDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &subjectTypeDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_type_dto",
				logValue: subjectTypeDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewSubjectTypeRouteGet
// @Summary Getting subjectTypes
// @Description Get all subjectTypes from database
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetSubjectTypesResponse}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes/ [get]
func NewSubjectTypeRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewSubjectTypeRouteGet"
	log = log.With(slog.String("op", op))

	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.GET("/", func(c *gin.Context) {
		resp, err := r.uc.Get(c)
		if err != nil {
			makeErrResponse(c, &ErrResp{
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

// NewSubjectTypeRouteGetByUUID
// @Summary Getting subjectType by uuid
// @Description Get subjectType from database with given uuid
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept */*
// @Produce json
// @Param uuid path string true "SubjectType uuid"
// @Success 200 {object} ResponseOK{response=models.SubjectType}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes/uuid/{uuid} [get]
func NewSubjectTypeRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_type_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewSubjectTypeRouteGetByType
// @Summary Getting subjectType by type
// @Description Get subjectType from database with given type
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept */*
// @Produce json
// @Param type path string true "Subject type"
// @Success 200 {object} ResponseOK{response=models.SubjectType}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes/type/{type} [get]
func NewSubjectTypeRouteGetByType(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.GET("/type/:type", func(c *gin.Context) {
		reqType := c.Param("type")

		resp, err := r.uc.GetByType(c, reqType)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_type",
				logValue: reqType,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewSubjectTypeRouteUpdate
// @Summary Updating subjectType
// @Description Update subjectType in database
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept json
// @Produce json
// @Param uuid path string true "SubjectType uuid"
// @Param subjectType body dto.UpdateSubjectTypeRequest true "SubjectType"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes/{uuid} [put]
func NewSubjectTypeRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.PUT("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var subjectTypeDTO dto.UpdateSubjectTypeRequest
		if err := c.ShouldBindJSON(&subjectTypeDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &subjectTypeDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_type",
				logValue: map[string]any{"uuid": reqUUID, "subject_type_dto": subjectTypeDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewSubjectTypeRouteDelete
// @Summary Deleting existing subjectType
// @Description Deleting existing subjectType from the database
// @Security ApiKeyAuth
// @Tags subjectType
// @Accept */*
// @Produce json
// @Param uuid path string true "SubjectType uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/subjecttypes/{uuid} [delete]
func NewSubjectTypeRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.SubjectTypeUseCase, log *slog.Logger) {
	r := &subjectTypeRoutes{uc, log}

	subjectTypeGroup := apiV1Group.Group("/subjecttypes")

	subjectTypeGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "subject_type_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
