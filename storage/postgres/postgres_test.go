package postgres_test

import (
	"os"
	"stad_projekt/config"
	"stad_projekt/storage"
	"stad_projekt/storage/postgres"
	"testing"

	_ "github.com/lib/pq"
)

var cfg config.Config = config.Config{
	PostgresHost:     "localhost",
	PostgresPort:     7052,
	PostgresUser:     "jasur",
	PostgresPassword: "123",
	PostgresDatabase: "pleace_db",
}

var strg storage.StorageI

func TestMain(m *testing.M) {
	pgfg := postgres.NewPostgres(cfg)
	strg = pgfg
	os.Exit(m.Run())
}
