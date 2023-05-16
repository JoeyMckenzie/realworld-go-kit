package main

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/internal/features"
	"golang.org/x/exp/slog"
)

func main() {
	// First, spin up our internal dependencies and logger
	logger := slog.Default()
	dataSourceName := os.Getenv("DSN")
	ctx := context.Background()

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

	// Initialize the service container and internal router
	logger.InfoCtx(ctx, "data connection successfully initialized, building initializing services")
	serviceContainer := features.NewServiceContainer(logger, db)

	internal.SeedDatabase(context.Background(), serviceContainer)
}
