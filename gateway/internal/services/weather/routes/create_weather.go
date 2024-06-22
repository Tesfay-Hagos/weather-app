package routes

import (
	"context"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/support"
	"github.com/gin-gonic/gin"
)

func CreateWeather(ctx *gin.Context, c pb.WeatherServiceClient, req models.CreateWeatherRequest) (models.CreateWeatherResponse, error) {
	res, err := c.CreateWeather(context.Background(), &pb.CreateWeatherRequest{
		City:      req.City,
		EndDate:   support.TimeToTimeStamp(req.EndDate),
		StartDate: support.TimeToTimeStamp(req.StartDate),
	})
	if err != nil {
		return models.CreateWeatherResponse{}, err
	}
	return *models.CopyFromPBToJson(res), nil
}
