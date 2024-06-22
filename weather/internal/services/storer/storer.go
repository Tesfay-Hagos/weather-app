package storer

import (
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storer interface {
	CreateWeather(models.Weather) error
}

type Weather struct {
	db *mongo.Database
}

func NewWeatherStorer(db *mongo.Database) *Weather {
	return &Weather{db: db}
}
func (w *Weather) CreateWeather(weather models.Weather) error {
	_, err := w.db.Collection("weather").InsertOne(nil, weather)
	if err != nil {
		return err
	}
	return nil
}
