package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type roomRoutes struct {
	uc  *usecase.RoomUseCase
	log *slog.Logger
}

// NewRoomRouteCreate
// @Summary Creating a new room
// @Description Creates a new room in the database and returns its uuid
// @Security ApiKeyAuth
// @Tags room
// @Accept json
// @Produce json
// @Param room body dto.CreateRoomRequest true "Room number"
// @Success 200 {object} ResponseOK{response=dto.CreateRoomRequest}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
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
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "room_dto",
				logValue: roomDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewRoomRouteGet
// @Summary Getting rooms
// @Description Get all rooms from database
// @Security ApiKeyAuth
// @Tags room
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetRoomsResponse}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
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

// NewRoomRouteGetByUUID
// @Summary Getting room by uuid
// @Description Get room from database with given uuid
// @Security ApiKeyAuth
// @Tags room
// @Accept */*
// @Produce json
// @Param uuid path string true "Room uuid"
// @Success 200 {object} ResponseOK{response=models.Room}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/uuid/{uuid} [get]
func NewRoomRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
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

// NewRoomRouteGetByNumber
// @Summary Getting room by number
// @Description Get room from database with given number
// @Security ApiKeyAuth
// @Tags room
// @Accept */*
// @Produce json
// @Param number path string true "Room number"
// @Success 200 {object} ResponseOK{response=models.Room}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
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

// NewRoomRouteUpdate
// @Summary Updating room
// @Description Update room in database
// @Security ApiKeyAuth
// @Tags room
// @Accept json
// @Produce json
// @Param uuid path string true "Room uuid"
// @Param room body dto.UpdateRoomRequest true "Room"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/{uuid} [put]
func NewRoomRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.PUT("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var roomDTO dto.UpdateRoomRequest
		if err := c.ShouldBindJSON(&roomDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &roomDTO)
		if err != nil {
			makeErrResponse(&ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "room",
				logValue: map[string]any{"uuid": reqUUID, "room_dto": roomDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewRoomRouteDelete
// @Summary Deleting existing room
// @Description Deleting existing room from the database
// @Security ApiKeyAuth
// @Tags room
// @Accept */*
// @Produce json
// @Param uuid path string true "Room uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/rooms/{uuid} [delete]
func NewRoomRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.RoomUseCase, log *slog.Logger) {
	r := &roomRoutes{uc, log}

	roomGroup := apiV1Group.Group("/rooms")

	roomGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
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

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
