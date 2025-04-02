package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"raspyx/config"
	_ "raspyx/docs"
	v1 "raspyx/internal/delivery/http"
	"raspyx/internal/delivery/http/middleware"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	gin.SetMode(gin.ReleaseMode)

	// Logger setup
	log, err := setupLogger(cfg)
	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("starting %v v%v", cfg.App.Name, cfg.App.Version), slog.String("logLevel", cfg.Log.Level))
	log.Debug("debug messages are enabled")

	// Creating db connection
	conn, err := pgx.Connect(context.Background(), cfg.PG.PGURL)
	if err != nil {
		log.Error("error db connection: %v", err)
		panic(err)
	}
	defer conn.Close(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.New()
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())

	// Pinger
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// All routes
	v1.NewRouter(r, log, conn)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.HTTP.Port),
		Handler: r,
	}

	log.Info(fmt.Sprintf("starting server at :%v", cfg.HTTP.Port))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("error starting server: %v", err))
			panic(err)
		}
	}()

	// shutdown
	<-ctx.Done()

	stop()
	log.Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", slog.String("error", err.Error()))
	}

	log.Info("server stopped")
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
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		return nil
	}
	return &lvl
}
