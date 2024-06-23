package weather

import (
	"net/http"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/routes"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/weather")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateWeather)
}

// CreateWeatherData
//
//	@Summary		Store weather data of your prefered city
//	@Description	Get and store weather data of your prefered city with the given time range.
//	@Tags			Weather-API
//	@Accept			json
//	@Produce		json
//	@param			weather		body		models.CreateWeatherRequest	true	"Create Weather Request body"
//	@Success		201			{object}	models.CreateWeatherResponse
//	@Failure		400,401,500	{object}	models.ErrorResponse
//	@Router			/weather [POST]
//
//	@Security		BearerAuth
func (svc *ServiceClient) CreateWeather(ctx *gin.Context) {
	b := models.CreateWeatherRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}
	if err := b.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}
	res, err := routes.CreateWeather(ctx, svc.Client, b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
