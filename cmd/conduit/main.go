package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joeymckenzie/realworld-go-kit/internal"
)

func main() {
	// First, spin up our internal dependencies and logger
	logger := internal.NewLogger()
	dataSourceName := os.Getenv("DSN")
	port := os.Getenv("PORT")
	parsedPost, err := strconv.Atoi(port)

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to parse port", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "initializing database connection...")

	// Grab a connection pool from the database
	db, err := sqlx.Open("mysql", dataSourceName)

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	// Run a quick ping check to make sure we're able to connect
	if err := db.Ping(); err != nil {
		level.Error(logger).Log("bootstrap", "failed to ping database", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "database connection successfully initialized, building routes")

	// Initialize the service container and internal router
	serviceContainer := internal.MakeServiceContainer(logger, db)
	router := internal.NewRouter(logger, serviceContainer)

	level.Info(logger).Log("bootstrap", fmt.Sprintf("routes successfully initialized, now listening on port %d", parsedPost))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", parsedPost), router); err != nil {
		level.Error(logger).Log("bootstrap", "failed to start server", "err", err)
		os.Exit(1)
	}
}
