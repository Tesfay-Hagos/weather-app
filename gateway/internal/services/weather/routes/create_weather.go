package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/pb"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/support"
	"github.com/gin-gonic/gin"
)

type CreateWeatherRequest struct {
	City      string    `json:"city"`
	endDate   time.Time `json:"endDate"`
	startDate time.Time `json:"startDate"`
}

func CreateWeather(ctx *gin.Context, c pb.WeatherServiceClient) {
	b := CreateWeatherRequest{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.CreateWeather(context.Background(), &pb.CreateWeatherRequest{
		City:      b.City,
		EndDate:   support.TimeToTimeStamp(b.endDate),
		StartDate: support.TimeToTimeStamp(b.startDate),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
