package models

import (
	"fmt"
	"time"

	"github.com/Tesfay-Hagos/go-grpc-api-gateway/internal/services/weather/constant/pb"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateWeatherRequest struct {
	City      string    `json:"city"`
	EndDate   time.Time `json:"endDate"`
	StartDate time.Time `json:"startDate"`
}

func (req CreateWeatherRequest) Validate() error {
	err := validation.ValidateStruct(req,
		// Validate email field
		validation.Field(&req.City,
			validation.Required.Error("email is required"),
			validation.Length(5, 100),
		),
		// Validate password field
		validation.Field(&req.EndDate,
			validation.Required.Error("password is required"),
		),
		// Validate confirm password field
		validation.Field(&req.StartDate,
			validation.Required.Error("confirm password is required"),
		),
	)

	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

type CreateWeatherResponse struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Address   string     `json:"address"`
	Timezone  string     `json:"timezone"`
	Days      []BasicDay `json:"days"`
}

type BasicDay struct {
	Datetime   string      `json:"datetime"`
	TempMax    float64     `json:"tempMax"`
	TempMin    float64     `json:"tempMin"`
	Conditions string      `json:"conditions"`
	PrecipProb float64     `json:"precipProb"`
	WindSpeed  float64     `json:"windSpeed"`
	WindDir    float64     `json:"windDir"`
	Humidity   float64     `json:"humidity"`
	Hours      []BasicHour `json:"hours"`
}

type BasicHour struct {
	Datetime   string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
	PrecipProb float64 `json:"precipProb"`
	WindSpeed  float64 `json:"windSpeed"`
	Humidity   float64 `json:"humidity"`
}

func CopyFromPBToJson(pbResp *pb.CreateWeatherResponse) *CreateWeatherResponse {
	jsonResp := &CreateWeatherResponse{
		Latitude:  pbResp.Latitude,
		Longitude: pbResp.Longitude,
		Address:   pbResp.Address,
		Timezone:  pbResp.Timezone,
		Days:      make([]BasicDay, len(pbResp.Days)),
	}

	for i, pbDay := range pbResp.Days {
		jsonDay := BasicDay{
			Datetime:   pbDay.Datetime,
			TempMax:    pbDay.TempMax,
			TempMin:    pbDay.TempMin,
			Conditions: pbDay.Conditions,
			PrecipProb: pbDay.PrecipProb,
			WindSpeed:  pbDay.WindSpeed,
			WindDir:    pbDay.WindDir,
			Humidity:   pbDay.Humidity,
			Hours:      make([]BasicHour, len(pbDay.Hours)),
		}

		for j, pbHour := range pbDay.Hours {
			jsonHour := BasicHour{
				Datetime:   pbHour.Datetime,
				Temp:       pbHour.Temp,
				Conditions: pbHour.Conditions,
				PrecipProb: pbHour.PrecipProb,
				WindSpeed:  pbHour.WindSpeed,
				Humidity:   pbHour.Humidity,
			}
			jsonDay.Hours[j] = jsonHour
		}

		jsonResp.Days[i] = jsonDay
	}

	return jsonResp
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
