package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/db"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/pb"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/services"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	run()
}
func run() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(log.Writer()), zapcore.InfoLevel))
	h := db.InitDB(c.DATABASEURL, logger)
	persistance := db.New(h, logger, db.Options{})
	m := db.InitiateMigration(c.MigrationPath, c.DATABASEURL, logger)
	db.UpMigration(m, logger)
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	fmt.Println("Auth Svc on", c.Port)
	s := services.Server{
		H:   persistance,
		Jwt: jwt,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
