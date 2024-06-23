package weather

import (
	"fmt"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/constant/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.WeatherServiceClient
}

func InitServiceClient(c *config.Config) pb.WeatherServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.NewClient(c.WeatherSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewWeatherServiceClient(cc)
}
