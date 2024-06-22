package createweather

import (
	"context"
	"testing"
	"time"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestCreateWeather(t *testing.T) {
	// Create a mock request
	req := &pb.CreateWeatherRequest{
		City:      "London",
		StartDate: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		EndDate:   &timestamp.Timestamp{Seconds: time.Now().AddDate(0, 0, 1).Unix()},
	}

	// Call the CreateWeather function
	res, err := s.CreateWeather(context.Background(), req)

	// Check if there was an error
	assert.NoError(t, err)

	// Check if the response is not nil
	assert.NotNil(t, res)
	assert.NotNil(t, res.Days)
	assert.NotNil(t, res.Days[0].Hours)
}
