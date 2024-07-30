package server

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/config"
	"github.com/cesc1802/onboarding-and-volunteer-service/database"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature"
	"log"
)

func RegisterServer() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	// Initialize router
	r := feature.RegisterHandlerV1(db)
	// Start server
	log.Printf("Server is running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
