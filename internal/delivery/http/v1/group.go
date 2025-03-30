package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
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

	// POST /api/v1/groups
	groupGroup.POST("/", func(c *gin.Context) {
		var groupDTO dto.CreateGroupRequest
		if err := c.ShouldBindJSON(&groupDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrWrongDataStructure})
			return
		}

		resp, err := r.uc.Create(c, &groupDTO)
		if err != nil {
			if strings.Contains(err.Error(), "group is not valid") {
				c.JSON(http.StatusBadRequest, RespError("Group is not valid"))
			} else if strings.Contains(err.Error(), "exist") {
				c.JSON(http.StatusBadRequest, RespError("Group exists"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
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
			c.JSON(http.StatusBadRequest, RespError("Invalid uuid"))
			return
		}

		err = r.uc.Delete(c, groupUUID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusBadRequest, RespError("Group not found"))
			} else {
				c.JSON(http.StatusInternalServerError, RespError("Internal server error"))
			}
			return
		}
		c.JSON(http.StatusOK, RespOK(nil))
	})
}
