package storer

import (
	"time"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeatherStor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	City      string             `bson:"city,omitempty"`
	StartDate time.Time          `bson:"startDate,omitempty"`
	EndDate   time.Time          `bson:"endDate,omitempty"`
	Weather   models.Weather     `bson:"weather,omitempty"`
}
