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

type groupRoutes struct {
	uc  *usecase.GroupUseCase
	log *slog.Logger
}

// NewGroupRouteCreate
// @Summary Creating a new group
// @Description Creates a new group in the database and returns its uuid
// @Tags group
// @Accept json
// @Produce json
// @Param group body dto.CreateGroupRequest true "Group number"
// @Success 200 {object} ResponseOK{response=dto.CreateGroupResponse}
// @Failure 400 {object} ResponseError
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
			if strings.Contains(err.Error(), "group is not valid") {
				log.Warn("Group is not valid", slog.String("group_number", groupDTO.Group))
				c.JSON(http.StatusBadRequest, RespError("Group is not valid"))
			} else if strings.Contains(err.Error(), "exist") {
				log.Info("Group exist", slog.String("group_number", groupDTO.Group))
				c.JSON(http.StatusBadRequest, RespError("Group exists"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewGroupRouteGet
// @Summary Getting groups
// @Description Get all groups from database
// @Tags group
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetGroupsResponse}
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups [get]
func NewGroupRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.GET("/", func(c *gin.Context) {
		resp, err := r.uc.Get(c)
		if err != nil {
			log.Error("Internal server error", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewGroupRouteGetByUUID
// @Summary Getting group by uuid
// @Description Get group from database with given uuid
// @Tags group
// @Accept */*
// @Produce json
// @Param uuid path string true "Group uuid"
// @Success 200 {object} ResponseOK{response=models.Group}
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/uuid/{uuid} [get]
func NewGroupRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		groupUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		resp, err := r.uc.GetByUUID(c, groupUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Group not found", slog.String("group_uuid", groupUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Group not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewGroupRouteGetByNumber
// @Summary Getting group by number
// @Description Get group from database with given number
// @Tags group
// @Accept */*
// @Produce json
// @Param number path string true "Group number"
// @Success 200 {object} ResponseOK{response=models.Group}
// @Failure 400 {object} ResponseError
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
			if strings.Contains(err.Error(), "not found") {
				log.Info("Group not found", slog.String("group_number", reqNumber))
				c.JSON(http.StatusNotFound, RespError("Group not found"))
			} else if strings.Contains(err.Error(), "group is not valid") {
				log.Warn("Group is not valid", slog.String("group_number", reqNumber))
				c.JSON(http.StatusBadRequest, RespError("Group is not valid"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewGroupRouteUpdate
// @Summary Updating group
// @Description Update group in database
// @Tags group
// @Accept json
// @Produce json
// @Param uuid path string true "Group uuid"
// @Param group body dto.UpdateGroupRequest true "Group"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/uuid/{uuid} [put]
func NewGroupRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.PUT("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		groupUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		var groupDTO dto.UpdateGroupRequest
		if err := c.ShouldBindJSON(&groupDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err = r.uc.Update(c, &models.Group{UUID: groupUUID, Number: groupDTO.Group})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Group not found", slog.String("group_uuid", groupUUID.String()))
				c.JSON(http.StatusNotFound, RespError("Group not found"))
			} else if strings.Contains(err.Error(), "group is not valid") {
				log.Warn("Group is not valid", slog.String("group_number", groupDTO.Group))
				c.JSON(http.StatusBadRequest, RespError("Group is not valid"))
			} else if strings.Contains(err.Error(), "exist") {
				log.Info("Group exist", slog.String("group_number", groupDTO.Group))
				c.JSON(http.StatusBadRequest, RespError("Group exists"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewGroupRouteDelete
// @Summary Deleting existing group
// @Description Deleting existing group from the database
// @Tags group
// @Accept */*
// @Produce json
// @Param uuid path string true "Group uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/groups/{uuid} [delete]
func NewGroupRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")

	groupGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		groupUUID, err := uuid.Parse(reqUUID)
		if err != nil {
			log.Warn("Invalid uuid", slog.String("error", err.Error()), slog.String("uuid", reqUUID))
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, groupUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Info("Group not found", slog.String("group_uuid", groupUUID.String()))
				c.JSON(http.StatusBadRequest, RespError("Group not found"))
			} else {
				log.Error("Internal server error", slog.String("error", err.Error()))
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
