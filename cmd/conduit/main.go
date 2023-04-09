package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/log/level"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeymckenzie/realworld-go-kit/internal"
)

func main() {
	// First, spin up our internal dependencies and logger
	logger := internal.NewLogger()
	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	ctx := context.Background()
	parsedPost, err := strconv.Atoi(port)

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to parse port", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "initializing database connection...")

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

	// Initialize the internal router and all services they'll be using
	router := internal.NewRouter(db)

	level.Info(logger).Log("bootstrap", fmt.Sprintf("routes successfully initialized, now listening on port %d", parsedPost))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", parsedPost), router); err != nil {
		level.Error(logger).Log("bootstrap", "failed to start server", "err", err)
		os.Exit(1)
	}
}
