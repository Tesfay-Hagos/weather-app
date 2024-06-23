package services

import (
	"context"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/models"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/services/storer"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/support"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	pb.UnimplementedWeatherServiceServer
	storer storer.Storer
	APIKey string
	URL    string
}

func NewServer(store storer.Storer, apikey, url string) *Server {
	return &Server{storer: store,
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
	err = s.storer.CreateWeather(ctx, storer.WeatherStor{
		ID:        primitive.NewObjectID(),
		City:      req.City,
		StartDate: req.StartDate.AsTime(),
		EndDate:   req.EndDate.AsTime(),
		Weather:   *we,
	})
	if err != nil {
		return &pb.CreateWeatherResponse{}, err
	}

	return models.ConvertWeatherToProto(*we), nil
}
