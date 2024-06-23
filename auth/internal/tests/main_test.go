package tests

import (
	"log"
	"testing"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/config"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/db"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/services"
	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var service services.Server

func TestMain(m *testing.M) {
	c, err := config.LoadConfig("../config/envs")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(log.Writer()), zapcore.InfoLevel))
	h := db.InitDB(c.DATABASEURL, logger)
	persistance := db.New(h, logger, db.Options{})
	mdb := db.InitiateMigration("../constant/db/schema", c.DATABASEURL, logger)
	db.UpMigration(mdb, logger)
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	service = *services.NewServer(persistance, jwt)
	m.Run()
}
