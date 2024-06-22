package services

import (
	"context"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/support"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedWeatherServiceServer
	db     *mongo.Database
	APIKey string
	URL    string
}

func NewServer(db *mongo.Database, apikey, url string) *Server {
	return &Server{db: db,
		APIKey: apikey,
		URL:    url}
}
func (s *Server) CreateWeather(ctx context.Context, req *pb.CreateWeatherRequest) (*pb.CreateWeatherResponse, error) {
	we, err := support.GetWeather(
		req.City,
		s.URL,
		s.APIKey,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return &pb.CreateWeatherResponse{}, err
	}

	return models.ConvertWeatherToProto(*we), nil
}
