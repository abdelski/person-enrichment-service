package main

import (
	"fmt"
	"log"
	"person-enrichment-service/pkg/http"
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
		log.Fatalf("Error loading config: %v", err)
	}




	server, err := http.NewServer(cfg)
	if err != nil {
		fmt.Println("Error initializing server:", err)
		return
	}


	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}