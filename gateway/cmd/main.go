package main

import (
	"log"

	_ "github.com/Tesfay-Hagos/go-grpc-api-gateway/docs"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Initiate
//
//	@title						Weather API Gateway service
//	@version					0.1.0
//	@description				This is API used to create weather data for any city in the world.
//
//	@contact.name				weather app support
//	@contact.url				http://www.weather.com
//	@contact.email				weather.support@weather.com
//
//	@host						localhost:8000
//	@BasePath					/v1
//	@termsOfService				http://www.weather.com/terms/
//
//	@license.name				License
//	@license.url				http://www.weather.com/license
//
//	@externalDocs.description	Find more info here
//	@externalDocs.url			https://github.com/weatherteam/api
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				"Bearer token"
//
//	@produce					application/json
func main() {
	run()
}
func run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Create a new gin engine
	router := gin.Default()
	group := router.Group("/v1")
	// Register the swagger handler
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Register routes for authentication, products, and orders
	authService := auth.RegisterRoutes(group, &cfg)
	weather.RegisterRoutes(group, &cfg, authService)

	// Run the gin engine
	router.Run(cfg.Port)
}
