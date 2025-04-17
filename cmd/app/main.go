package main

import (
	"log"
	"raspyx/config"
	_ "raspyx/docs"
	"raspyx/internal/app"
)

// @title           Raspyx
// @version         1.1.0
// @description     API for schedules

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
