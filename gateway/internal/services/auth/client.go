package auth

import (
	"fmt"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.NewClient(c.AuthSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
