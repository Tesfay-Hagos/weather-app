package storer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeatherStruct struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	City      string             `bson:"city,omitempty"`
	StartDate time.Time          `bson:"startDate,omitempty"`
	EndDate   time.Time          `bson:"endDate,omitempty"`
	Weather   string             `bson:"weather,omitempty"`
}
