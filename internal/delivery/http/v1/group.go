package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type groupRoutes struct {
	uc  *usecase.GroupUseCase
	log *slog.Logger
}

// NewGroupRouteCreate
// @Summary Creating a new group
// @Description Creates a new group in the database and returns its uuid
// @Security ApiKeyAuth
// @Tags group
// @Accept json
// @Produce json
// @Param group body dto.CreateGroupRequest true "Group number"
// @Success 200 {object} ResponseOK{response=dto.CreateGroupResponse}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups [post]
func NewGroupRouteCreate(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.POST("/", func(c *gin.Context) {
		var groupDTO dto.CreateGroupRequest
		if err := c.ShouldBindJSON(&groupDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &groupDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "group_dto",
				logValue: groupDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewGroupRouteGet
// @Summary Getting groups
// @Description Get all groups from database
// @Security ApiKeyAuth
// @Security ApiKeyAuth
// @Tags group
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetGroupsResponse}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/ [get]
func NewGroupRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.GET("/", func(c *gin.Context) {
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

// NewGroupRouteGetByUUID
// @Summary Getting group by uuid
// @Description Get group from database with given uuid
// @Security ApiKeyAuth
// @Tags group
// @Accept */*
// @Produce json
// @Param uuid path string true "Group uuid"
// @Success 200 {object} ResponseOK{response=models.Group}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/uuid/{uuid} [get]
func NewGroupRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
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

// NewGroupRouteGetByNumber
// @Summary Getting group by number
// @Description Get group from database with given number
// @Security ApiKeyAuth
// @Tags group
// @Accept */*
// @Produce json
// @Param number path string true "Group number"
// @Success 200 {object} ResponseOK{response=models.Group}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/number/{number} [get]
func NewGroupRouteGetByNumber(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.GET("/number/:number", func(c *gin.Context) {
		reqNumber := c.Param("number")
		resp, err := r.uc.GetByNumber(c, reqNumber)
		if err != nil {
			makeErrResponse(c, &ErrResp{
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

// NewGroupRouteUpdate
// @Summary Updating group
// @Description Update group in database
// @Security ApiKeyAuth
// @Tags group
// @Accept json
// @Produce json
// @Param uuid path string true "Group uuid"
// @Param group body dto.UpdateGroupRequest true "Group"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/{uuid} [put]
func NewGroupRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.PUT("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var groupDTO dto.UpdateGroupRequest
		if err := c.ShouldBindJSON(&groupDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &groupDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "group",
				logValue: map[string]any{"uuid": reqUUID, "group_dto": groupDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewGroupRouteDelete
// @Summary Deleting existing group
// @Description Deleting existing group from the database
// @Security ApiKeyAuth
// @Tags group
// @Accept */*
// @Produce json
// @Param uuid path string true "Group uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/{uuid} [delete]
func NewGroupRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "group_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
