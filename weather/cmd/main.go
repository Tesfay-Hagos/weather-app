package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/services"
	"github.com/Tesfay-Hagos/go-grpc-weather-svc/internal/services/storer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	run()
}

func run() {
	c, err := config.LoadConfig("./internal/config/envs")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.DATABASEURL))
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)
	storer := storer.NewWeatherStorer(client.Database(c.DbName))
	s := services.NewServer(storer, c.WeatherAPIKey, c.WeatherURL)

	grpcServer := grpc.NewServer()

	pb.RegisterWeatherServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
