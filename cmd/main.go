package main

import (
	"log"
	"person-enrichment-service/pkg/http"
	"person-enrichment-service/pkg/logging"
	"person-enrichment-service/server/config"
)

// @title Person Enrichment API
// @version 1.0
// @description This is a service for enriching person data with age, gender, and nationality
// @host localhost:8081
// @BasePath /api/v1
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: %v", err)
	}

	logger := logging.NewLogger(cfg.LogLevel)

	server, err := http.NewServer(cfg, logger)
	if err != nil {
		logger.Error("Error creating server: %v", err)
		return
	}

	if err := server.Start(); err != nil {
		logger.Error("Error starting server: %v", err)
		return
	}
}
