package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type locationRoutes struct {
	uc  *usecase.LocationUseCase
	log *slog.Logger
}

// NewLocationRouteCreate
// @Summary Creating a new location
// @Description Creates a new location in the database and returns its uuid
// @Tags location
// @Accept json
// @Produce json
// @Param location body dto.CreateLocationRequest true "Location name"
// @Success 200 {object} ResponseOK{response=dto.CreateLocationResponse}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations [post]
func NewLocationRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.POST("/", func(c *gin.Context) {
		var locationDTO dto.CreateLocationRequest
		if err := c.ShouldBindJSON(&locationDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &locationDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "location_dto",
				logValue: locationDTO,
			})
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewLocationRouteGet
// @Summary Getting locations
// @Description Get all locations from database
// @Tags location
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetLocationsResponse}
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations [get]
func NewLocationRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.GET("/", func(c *gin.Context) {
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

// NewLocationRouteGetByUUID
// @Summary Getting location by uuid
// @Description Get location from database with given uuid
// @Tags location
// @Accept */*
// @Produce json
// @Param uuid path string true "Location uuid"
// @Success 200 {object} ResponseOK{response=models.Location}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations/uuid/{uuid} [get]
func NewLocationRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
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

// NewLocationRouteGetByName
// @Summary Getting location by name
// @Description Get location from database with given name
// @Tags location
// @Accept */*
// @Produce json
// @Param name path string true "Location name"
// @Success 200 {object} ResponseOK{response=models.Location}
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations/name/{name} [get]
func NewLocationRouteGetByName(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.GET("/name/:name", func(c *gin.Context) {
		reqName := c.Param("name")
		resp, err := r.uc.GetByName(c, reqName)
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

// NewLocationRouteUpdate
// @Summary Updating location
// @Description Update location in database
// @Tags location
// @Accept json
// @Produce json
// @Param uuid path string true "Location uuid"
// @Param location body dto.UpdateLocationRequest true "Location"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations/uuid/{uuid} [put]
func NewLocationRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var locationDTO dto.UpdateLocationRequest
		if err := c.ShouldBindJSON(&locationDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &locationDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "location",
				logValue: map[string]any{"uuid": reqUUID, "location_dto": locationDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewLocationRouteDelete
// @Summary Deleting existing location
// @Description Deleting existing location from the database
// @Tags location
// @Accept */*
// @Produce json
// @Param uuid path string true "Location uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations/{uuid} [delete]
func NewLocationRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
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

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
