package support

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/models"
	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func GetWeather(location, url, apiKey string, start, end time.Time) (*models.Weather, error) {
	startDate := start.Format("2006-01-01")
	endDate := end.Format("2006-01-03")

	client := resty.New()

	resp, err := client.R().
		SetQueryParam("key", apiKey).
		SetQueryParam("location", location).
		SetQueryParam("startDate", startDate).
		SetQueryParam("endDate", endDate).
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("Error: %v", resp)
		return nil, err
	}

	weatherData := &models.Weather{}

	err = json.Unmarshal(resp.Body(), weatherData)
	return weatherData, err
}

func TimeStampToTime(t *timestamp.Timestamp) time.Time {
	return t.AsTime()
}
