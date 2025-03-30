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
				c.JSON(http.StatusBadRequest, gin.H{"error": "Group is not valid"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
