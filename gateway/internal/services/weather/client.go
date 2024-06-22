package weather

import (
	"fmt"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.WeatherServiceClient
}

func InitServiceClient(c *config.Config) pb.WeatherServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.WeatherSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewWeatherServiceClient(cc)
}
