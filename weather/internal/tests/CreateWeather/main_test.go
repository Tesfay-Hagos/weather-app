package createweather

import (
	"context"
	"log"
	"testing"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/services"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/services/storer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var s *services.Server

func TestMain(m *testing.M) {
	// setup
	c, err := config.LoadConfig("../../config/envs")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.DATABASEURL))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	storer := storer.NewWeatherStorer(client.Database(c.DbName))
	s = services.NewServer(storer, c.WeatherAPIKey, c.WeatherURL)
	m.Run()
}
