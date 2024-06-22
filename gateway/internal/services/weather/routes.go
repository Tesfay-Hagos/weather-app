package weather

import (
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/routes"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/weather")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateWeather)
}

func (svc *ServiceClient) CreateWeather(ctx *gin.Context) {
	routes.CreateWeather(ctx, svc.Client)
}
