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
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &locationDTO)
		if err != nil {
			if strings.Contains(err.Error(), "exist") {
				c.JSON(http.StatusBadRequest, RespError("Location exists"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
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
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
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
		locationUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		resp, err := r.uc.GetByUUID(c, locationUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, RespError("Location not found"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
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
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, RespError("Location not found"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
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
		locationUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		var locationDTO dto.UpdateLocationRequest
		if err := c.ShouldBindJSON(&locationDTO); err != nil {
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err = r.uc.Update(c, &models.Location{UUID: locationUUID, Name: locationDTO.Name})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, RespError("Location not found"))
			} else if strings.Contains(err.Error(), "exist") {
				c.JSON(http.StatusBadRequest, RespError("Location exists"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
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
// @Failure 500 {object} ResponseError
// @Router /api/v1/locations/{uuid} [delete]
func NewLocationRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.LocationUseCase, log *slog.Logger) {
	r := &locationRoutes{uc, log}

	locationGroup := apiV1Group.Group("/locations")

	locationGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		locationUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, locationUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusBadRequest, RespError("Location not found"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}
		
		c.JSON(http.StatusOK, RespOK(nil))
	})
}
