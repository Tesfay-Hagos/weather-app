package routes

import (
	"context"
	"fmt"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/pb"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient, req models.RegisterRequestBody) (models.RegisterResponse, error) {
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if res == nil {
		return models.RegisterResponse{}, fmt.Errorf("failed to register user: %v", err.Error())
	}
	return models.PbRegisterResponseToRegisterResponse(res), err

}
