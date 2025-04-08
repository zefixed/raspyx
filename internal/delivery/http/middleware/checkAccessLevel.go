package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "raspyx/internal/delivery/http/v1"
)

func AccessLevelMiddleware(accessLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAccessLevel := int(c.GetFloat64("access_level"))
		if userAccessLevel < accessLevel {
			c.AbortWithStatusJSON(http.StatusForbidden, v1.RespError("forbidden"))
			return
		}
		c.Next()
	}
}
