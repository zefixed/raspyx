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

type roomRoutes struct {
	uc  *usecase.RoomUseCase
	log *slog.Logger
}

// NewRoomRouteCreate
// @Summary Creating a new room
// @Description Creates a new room in the database and returns its uuid
// @Tags room
// @Accept json
// @Produce json
// @Param room body dto.CreateRoomRequest true "Room number"
// @Success 200 {object} ResponseOK{response=dto.CreateRoomRequest}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms [post]
func NewRoomRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewRoomRouteCreate"
	log = log.With(slog.String("op", op))

	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.POST("/", func(c *gin.Context) {
		var roomDTO dto.CreateRoomRequest
		if err := c.ShouldBindJSON(&roomDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &roomDTO)
		if err != nil {
			if strings.Contains(err.Error(), "exist") {
				log.Info("Room exist", slog.String("room_number", roomDTO.Number))
				c.JSON(http.StatusBadRequest, RespError("Room exists"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewRoomRouteGet
// @Summary Getting rooms
// @Description Get all rooms from database
// @Tags room
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetRoomsResponse}
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms [get]
func NewRoomRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewRoomRouteGet"
	log = log.With(slog.String("op", op))

	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.GET("/", func(c *gin.Context) {
		resp, err := r.uc.Get(c)
		if err != nil {
			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewRoomRouteGetByUUID
// @Summary Getting room by uuid
// @Description Get room from database with given uuid
// @Tags room
// @Accept */*
// @Produce json
// @Param uuid path string true "Room uuid"
// @Success 200 {object} ResponseOK{response=models.Room}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/uuid/{uuid} [get]
func NewRoomRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		roomUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		resp, err := r.uc.GetByUUID(c, roomUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Room not found", slog.String("room_uuid", roomUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Room not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewRoomRouteGetByNumber
// @Summary Getting room by number
// @Description Get room from database with given number
// @Tags room
// @Accept */*
// @Produce json
// @Param number path string true "Room number"
// @Success 200 {object} ResponseOK{response=models.Room}
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/number/{number} [get]
func NewRoomRouteGetByNumber(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.GET("/number/:number", func(c *gin.Context) {
		reqNumber := c.Param("number")

		resp, err := r.uc.GetByNumber(c, reqNumber)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Room not found", slog.String("room_number", reqNumber))
				c.JSON(http.StatusNotFound, RespError("Room not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewRoomRouteUpdate
// @Summary Updating room
// @Description Update room in database
// @Tags room
// @Accept json
// @Produce json
// @Param uuid path string true "Room uuid"
// @Param room body dto.UpdateRoomRequest true "Room"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/uuid/{uuid} [put]
func NewRoomRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		roomUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		var roomDTO dto.UpdateRoomRequest
		if err := c.ShouldBindJSON(&roomDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err = r.uc.Update(c, &models.Room{UUID: roomUUID, Number: roomDTO.Number})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Room not found", slog.String("room_uuid", roomUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Room not found"))
			} else if strings.Contains(err.Error(), "exist") {
				log.Info("Room exist", slog.String("room_number", roomDTO.Number))
				c.JSON(http.StatusBadRequest, RespError("Room exists"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewRoomRouteDelete
// @Summary Deleting existing room
// @Description Deleting existing room from the database
// @Tags room
// @Accept */*
// @Produce json
// @Param uuid path string true "Room uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/{uuid} [delete]
func NewRoomRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		roomUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, roomUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Room not found", slog.String("room_uuid", roomUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Room not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
