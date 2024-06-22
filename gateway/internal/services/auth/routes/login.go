package routes

import (
	"context"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/pb"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient, req models.LoginRequestBody) (models.LoginResponse, error) {
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	return models.PbLoginResponseToLoginResponse(res), err

}
