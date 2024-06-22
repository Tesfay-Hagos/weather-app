package models

import "github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/pb"

// BasicWeather represents the overall basic weather forecast data
type Weather struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Address   string     `json:"address"`
	Timezone  string     `json:"timezone"`
	Days      []BasicDay `json:"days"`
}

// BasicDay represents the essential weather data for a specific day
type BasicDay struct {
	DateTime   string      `json:"datetime"`
	TempMax    float64     `json:"tempmax"`
	TempMin    float64     `json:"tempmin"`
	Conditions string      `json:"conditions"`
	PrecipProb float64     `json:"precipprob"`
	WindSpeed  float64     `json:"windspeed"`
	WindDir    float64     `json:"winddir"`
	Humidity   float64     `json:"humidity"`
	Hours      []BasicHour `json:"hours"`
}

// BasicHour represents the essential weather data for a specific hour
type BasicHour struct {
	DateTime   string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
	PrecipProb float64 `json:"precipprob"`
	WindSpeed  float64 `json:"windspeed"`
	Humidity   float64 `json:"humidity"`
}

// ConvertWeatherToProto converts a Weather struct to a CreateWeatherResponse protobuf message
func ConvertWeatherToProto(weather Weather) *pb.CreateWeatherResponse {
	var days []*pb.BasicDay
	for _, day := range weather.Days {
		var hours []*pb.BasicHour
		for _, hour := range day.Hours {
			protoHour := &pb.BasicHour{
				Datetime:   hour.DateTime,
				Temp:       hour.Temp,
				Conditions: hour.Conditions,
				PrecipProb: hour.PrecipProb,
				WindSpeed:  hour.WindSpeed,
				Humidity:   hour.Humidity,
			}
			hours = append(hours, protoHour)
		}

		protoDay := &pb.BasicDay{
			Datetime:   day.DateTime,
			TempMax:    day.TempMax,
			TempMin:    day.TempMin,
			Conditions: day.Conditions,
			PrecipProb: day.PrecipProb,
			WindSpeed:  day.WindSpeed,
			WindDir:    day.WindDir,
			Humidity:   day.Humidity,
			Hours:      hours,
		}
		days = append(days, protoDay)
	}

	return &pb.CreateWeatherResponse{
		Latitude:  weather.Latitude,
		Longitude: weather.Longitude,
		Address:   weather.Address,
		Timezone:  weather.Timezone,
		Days:      days,
	}
}
