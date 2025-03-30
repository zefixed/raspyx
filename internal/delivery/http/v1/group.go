package v1

import (
	"github.com/gin-gonic/gin"
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

// NewGroupRoutes
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
func NewGroupRoutes(apiV1Group *gin.RouterGroup, uc *usecase.GroupUseCase, log *slog.Logger) {
	r := &groupRoutes{uc, log}

	groupGroup := apiV1Group.Group("/groups")
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
