package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"raspyx/config"
)

func Run(cfg *config.Config) {
	log, err := setupLogger(cfg)
	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("starting %v v%v", cfg.App.Name, cfg.App.Version), slog.String("logLevel", cfg.Log.Level))
	log.Debug("debug messages are enabled")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(fmt.Sprintf(":%v", cfg.HTTP.Port))
	if err != nil {
		log.Error(fmt.Sprintf("error starting server, %v", err))
	}
}

func setupLogger(cfg *config.Config) (*slog.Logger, error) {
	var log *slog.Logger
	var err error

	var handler slog.Handler
	level := getLogLevel(cfg.Log.Level)

	if level == nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL=%v", cfg.Log.Level)
	}

	switch cfg.Log.Type {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: *level})
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: *level})
	default:
		return nil, fmt.Errorf("invalid LOG_TYPE=%v", cfg.Log.Type)
	}

	log = slog.New(handler)
	return log, err
}

func getLogLevel(level string) *slog.Level {
	var lvl slog.Level
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	default:
		return nil
	}
	return &lvl
}
