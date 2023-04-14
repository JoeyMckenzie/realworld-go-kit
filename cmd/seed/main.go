package main

import (
	"context"
	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/internal/features"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	const loggingSpan string = "seed"

	// First, spin up our internal dependencies and logger
	logger := internal.NewLogger()

	if err := godotenv.Load(); err != nil {
		level.Error(logger).Log(loggingSpan, "failed to local environment variables", "err", err)
		os.Exit(1)
	}

	dataSourceName := os.Getenv("DSN")

	// Grab a connection and verify we're able to ping PlanetScale
	level.Info(logger).Log(loggingSpan, "initializing data connection...")
	db, err := sqlx.Open("mysql", dataSourceName)

	if err != nil {
		level.Error(logger).Log(loggingSpan, "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		level.Error(logger).Log(loggingSpan, "failed to ping data", "err", err)
		os.Exit(1)
	}

	// Initialize the service container and internal router
	level.Info(logger).Log(loggingSpan, "data connection successfully initialized, building initializing services")
	serviceContainer := features.NewServiceContainer(logger, db)

	internal.SeedDatabase(context.Background(), serviceContainer)
}
