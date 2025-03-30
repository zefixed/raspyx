package main

import (
	"log"
	"raspyx/config"
	_ "raspyx/docs"
	"raspyx/internal/app"
)

// @title           Raspyx
// @version         0.0.1
// @description     API for schedules

// @host      localhost:8080
// @BasePath  /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
