package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// before request

		c.Next()

		// after request
		log.Info("request",
			slog.String("component", "middleware/logger"),
			slog.String("method", fmt.Sprintf("%v", c.Request.Method)),
			slog.String("status", fmt.Sprintf("%v", c.Writer.Status())),
			slog.String("url", fmt.Sprintf("%v", c.Request.URL.Path)),
			slog.String("client_ip", fmt.Sprintf("%v", c.ClientIP())),
			slog.String("user_agent", fmt.Sprintf("%v", c.Request.UserAgent())),
			slog.String("errors", fmt.Sprintf("%v", c.Errors.String())),
			slog.String("latency", fmt.Sprintf("%v", time.Since(startTime))))
	}
}
