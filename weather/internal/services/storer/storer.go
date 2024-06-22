package storer

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storer interface {
	CreateWeather(context.Context, WeatherStor) error
}

type Weather struct {
	db *mongo.Database
}

func NewWeatherStorer(db *mongo.Database) *Weather {
	return &Weather{db: db}
}
func (w *Weather) CreateWeather(ctx context.Context, weather WeatherStor) error {
	_, err := w.db.Collection("weather").InsertOne(ctx, weather)
	if err != nil {
		return err
	}
	return nil
}
