package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Tesfay-Hagos/go-grpc-auth-svc/internal/constant/db/queries/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // go-migrate needs it
	_ "github.com/golang-migrate/migrate/v4/source/file"       // go-migrate needs it
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitDB(url string, log *zap.Logger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	idleConnTimeout := viper.GetDuration("database.idle_conn_timeout")
	if idleConnTimeout == 0 {
		idleConnTimeout = 4 * time.Minute
	}
	config.MaxConnIdleTime = idleConnTimeout
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	if _, err := conn.Exec(context.Background(), "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'"); err != nil {
		log.Fatal(fmt.Sprintf("Failed to ping database: %v", err))
	}

	return conn
}
func InitiateMigration(path, conn string, log *zap.Logger) *migrate.Migrate {
	conn = fmt.Sprintf("postgres://%s", strings.Split(conn, "://")[1])
	m, err := migrate.New(fmt.Sprintf("file://%s", path), conn)
	if err != nil {
		log.Fatal("could not create migrator", zap.Error(err))
	}
	return m
}

func UpMigration(m *migrate.Migrate, log *zap.Logger) {
	err := m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("could not migrate", zap.Error(err))
	}
}

type PersistenceDB struct {
	*db.Queries
	pool    *pgxpool.Pool
	log     *zap.Logger
	options Options
}

type Options struct {
	SSODB        Sibling
	AuthzDB      Sibling
	AccountingDB Sibling
}

func setOptions(options Options) Options {
	if len(options.SSODB) == 0 {
		options.SSODB = "sso"
	}

	if len(options.AuthzDB) == 0 {
		options.AuthzDB = "authz"
	}

	if len(options.AccountingDB) == 0 {
		options.AccountingDB = "accounting"
	}

	return options
}

type Sibling string

func New(pool *pgxpool.Pool, log *zap.Logger, options Options) PersistenceDB {
	return PersistenceDB{
		Queries: db.New(pool),
		pool:    pool,
		log:     log,
		options: setOptions(options),
	}
}
