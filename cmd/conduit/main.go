package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/features"
	"golang.org/x/exp/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joeymckenzie/realworld-go-kit/internal"
)

func main() {
	// First, spin up our internal dependencies and logger
	logger := slog.Default()
	dataSourceName := os.Getenv("DSN")
	port := os.Getenv("PORT")
	parsedPost, err := strconv.Atoi(port)
	ctx := context.Background()

	if err != nil {
		logger.ErrorCtx(ctx, "failed to parse port", "err", err)
		os.Exit(1)
	}

	// Grab a connection and verify we're able to ping PlanetScale
	logger.InfoCtx(ctx, "initializing data connection...")
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		logger.ErrorCtx(ctx, "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		logger.ErrorCtx(ctx, "failed to ping data", "err", err)
		os.Exit(1)
	}

	logger.InfoCtx(ctx, "data connection successfully initialized, building routes")

	// Initialize the service container and internal router
	serviceContainer := features.NewServiceContainer(logger, db)
	router := internal.NewRouter(logger, serviceContainer)

	logger.InfoCtx(ctx, fmt.Sprintf("routes successfully initialized, now listening on port %d", parsedPost))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", parsedPost), router); err != nil {
		logger.ErrorCtx(ctx, "failed to start server", "err", err)
		os.Exit(1)
	}
}
