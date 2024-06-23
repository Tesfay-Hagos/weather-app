package tests

import (
	"context"
	"math/rand"
	"net/http"
	"strings"
	"testing"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/pb"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	// Create a valid RegisterRequest
	req := &pb.RegisterRequest{
		Email:           RandomEmailGenerator(),
		Password:        "John12@doe",
		ConfirmPassword: "John12@doe",
	}

	// Call the Register function
	res, err := service.Register(context.Background(), req)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the response status is http.StatusCreated
	assert.Equal(t, http.StatusCreated, int(res.Status))

	// Assert that the response UserId is not empty
	assert.NotEmpty(t, res.UserId)
}

func TestRegister_InvalidRequest(t *testing.T) {
	// Create an invalid RegisterRequest with an empty email
	req := &pb.RegisterRequest{
		Email:    "",
		Password: "password",
	}

	// Call the Register function
	res, err := service.Register(context.Background(), req)

	// Assert that the error is not nil
	assert.Error(t, err)

	// Assert that the response status is http.StatusBadRequest
	assert.Equal(t, http.StatusBadRequest, int(res.Status))

	// Assert that the response Error is not empty
	assert.NotEmpty(t, res.Error)
}

func TestRegister_ErrorCreatingUser(t *testing.T) {

	// Create a RegisterRequest
	req := &pb.RegisterRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	// Call the Register function with an error
	res, err := service.Register(context.Background(), req)

	// Assert that the error is not nil
	assert.Error(t, err)

	// Assert that the response status is http.StatusBadRequest
	assert.Equal(t, http.StatusBadRequest, int(res.Status))

	// Assert that the response Error is not empty
	assert.NotEmpty(t, res.Error)
}

func RandomEmailGenerator() string {
	return "test" + RandomStringGenerator(5) + "@example.com"

}

func RandomStringGenerator(length int) string {
	var result strings.Builder
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		result.WriteByte(characters[rand.Intn(len(characters))])
	}
	return result.String()
}
