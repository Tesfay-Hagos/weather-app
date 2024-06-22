package main

import (
	"log"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Create a new gin engine
	router := gin.Default()
	// Register routes for authentication, products, and orders
	authService := auth.RegisterRoutes(router, &cfg)
	weather.RegisterRoutes(router, &cfg, authService)

	// Run the gin engine
	router.Run(cfg.Port)
}
