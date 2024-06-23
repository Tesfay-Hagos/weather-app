package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/pb"
	"github.com/stretchr/testify/assert"
)

func TestLogin_InvalidRequest(t *testing.T) {
	// Create a valid RegisterRequest
	r := &pb.RegisterRequest{
		Email:           RandomEmailGenerator(),
		Password:        "John12@doe",
		ConfirmPassword: "John12@doe",
	}

	// Call the Register function
	_, err := service.Register(context.Background(), r)
	assert.NoError(t, err)
	// Create a valid LoginRequest
	req := &pb.LoginRequest{
		Email:    r.Email,
		Password: "123456",
	}

	// Call the Login function
	res, err := service.Login(context.Background(), req)

	// Assert that the error is not nil
	assert.Error(t, err)

	// Assert that the response status is http.StatusBadRequest
	assert.Equal(t, http.StatusNotFound, int(res.Status))

	// Assert that the response Error is not empty
	assert.NotEmpty(t, res.Error)
}

func TestLogin_ValidRequest(t *testing.T) {
	// Create a valid RegisterRequest
	r := &pb.RegisterRequest{
		Email:           RandomEmailGenerator(),
		Password:        "John12@doe",
		ConfirmPassword: "John12@doe",
	}

	// Call the Register function
	_, err := service.Register(context.Background(), r)
	assert.NoError(t, err)
	// Create a valid LoginRequest
	req := &pb.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	}

	// Call the Login function
	res, err := service.Login(context.Background(), req)

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the response status is http.StatusOK
	assert.Equal(t, http.StatusOK, int(res.Status))

	// Assert that the response Error is empty
	assert.Empty(t, res.Error)

	// Assert that the response Token is not empty
	assert.NotEmpty(t, res.Token)
}
