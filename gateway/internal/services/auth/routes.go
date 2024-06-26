package auth

import (
	"net/http"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

// UserRegistration
//
//	@Summary		Register new user
//	@Description	Register new user to the system.
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@param			register	body		models.RegisterRequestBody	true	"Register to the system"
//	@Success		201			{object}	models.RegisterResponse
//	@Failure		400,401,500	{object}	models.RegisterResponse
//	@Router			/auth/register [POST]
func (svc *ServiceClient) Register(ctx *gin.Context) {
	body := models.RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := body.Validate(); err != nil {
		errBody := &models.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errBody)
		return
	}
	res, err := routes.Register(ctx, svc.Client, body)
	if err != nil {
		errBody := &models.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errBody)
		return
	}
	ctx.JSON(int(res.Status), &res)
}

// UserLogin
//
//	@Summary		Login user
//	@Description	Login user to the system.
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@param			login		body		models.LoginRequestBody	true	"Login to the system"
//	@Success		201			{object}	models.LoginResponse
//	@Failure		400,401,500	{object}	models.LoginResponse
//	@Router			/auth/login [POST]
func (svc *ServiceClient) Login(ctx *gin.Context) {
	b := models.LoginRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		errorbody := models.LoginResponse{
			Status: 400,
			Error:  err.Error(),
			Token:  "",
		}
		ctx.JSON(http.StatusBadRequest, errorbody)
		return
	}
	if err := b.Validate(); err != nil {
		errorbody := models.LoginResponse{
			Status: 400,
			Error:  err.Error(),
			Token:  "",
		}
		ctx.JSON(http.StatusBadRequest, &errorbody)
		return
	}
	res, err := routes.Login(ctx, svc.Client, b)
	if err != nil {
		errorbody := models.LoginResponse{
			Status: 400,
			Error:  err.Error(),
			Token:  "",
		}
		ctx.JSON(http.StatusBadRequest, &errorbody)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
