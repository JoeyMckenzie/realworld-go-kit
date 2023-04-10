package main

import (
	"context"
	"os"

	"github.com/go-kit/log/level"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeymckenzie/realworld-go-kit/internal"
)

func main() {
	// First, spin up our internal dependencies and logger
	logger := internal.NewLogger()
	connectionString := os.Getenv("DATABASE_URL")
	ctx := context.Background()

	level.Info(logger).Log("seed", "initializing database connection...")

	// Grab a connection pool from the database
	db, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	// Run a quick ping check to make sure we're able to connect
	if err := db.Ping(ctx); err != nil {
		level.Error(logger).Log("bootstrap", "failed to ping database", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "database connection successfully initialized, building routes")
}
