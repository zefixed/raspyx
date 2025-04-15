package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"raspyx/config"
	_ "raspyx/docs"
	v1 "raspyx/internal/delivery/http"
	mw "raspyx/internal/delivery/http/middleware"
	"raspyx/internal/parser"
	"strings"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	gin.SetMode(gin.ReleaseMode)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Logger setup
	log, err := setupLogger(cfg)
	if err != nil {
		log.Error(fmt.Sprintf("error setting up loger: %v", err))
		return
	}

	log.Info(fmt.Sprintf("starting %v v%v", cfg.App.Name, cfg.App.Version), slog.String("logLevel", cfg.Log.Level))
	log.Debug("debug messages are enabled")

	// db connection
	conn, err := dbConn(ctx, cfg)
	if err != nil {
		log.Error(fmt.Sprintf("error db connection: %v", err))
		return
	}
	defer conn.Close(ctx)

	// redis client
	redisClient, err := cacheClient(ctx, cfg)
	if err != nil {
		log.Error(fmt.Sprintf("error redis cache: %v", err))
		return
	}
	defer redisClient.Close()

	// Router
	r := gin.New()
	r.Use(mw.Logger(log))
	r.Use(gin.Recovery())
	r.Use(mw.RequestIDMiddleware())

	// All routes
	v1.NewRouter(r, log, conn, redisClient, cfg)

	// Pinger
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.HTTP.Port),
		Handler: r,
	}

	log.Info(fmt.Sprintf("starting server at :%v", cfg.HTTP.Port))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("error starting server: %v", err))
			return
		}
	}()

	// Schedule parser
	parser.NewScheduleParser(10*time.Second, conn, log, cfg.Parser).New(ctx)

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

func dbConn(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	// Creating db connection
	conn, err := pgx.Connect(ctx, cfg.PG.PGURL)
	if err != nil {
		return nil, err
	}

	// Ping db connection
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func cacheClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	// Parsing redis url from config
	opt, err := redis.ParseURL(cfg.Redis.REDIS_URL)
	if err != nil {
		return nil, err
	}

	// Creating new redis client and ping it
	redisClient := redis.NewClient(opt)
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return redisClient, nil
}

func setupLogger(cfg *config.Config) (*slog.Logger, error) {
	var log *slog.Logger
	var err error

	var handler slog.Handler
	level := getLogLevel(strings.TrimSpace(cfg.Log.Level))

	if level == nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL=%v", cfg.Log.Level)
	}

	switch strings.TrimSpace(cfg.Log.Type) {
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
