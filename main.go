package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jackc/pgx/v5"
	"github.com/joeymckenzie/realworld-go-kit/app/users"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	parsedPost, err := strconv.Atoi(port)
	ctx := context.Background()

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to parse port", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "initializing database connection...")

	db, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		level.Error(logger).Log("bootstrap", "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	if err := db.Ping(ctx); err != nil {
		level.Error(logger).Log("bootstrap", "failed to ping database", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("bootstrap", "database connection successfully initialized, building routes")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	usersRepository := users.NewRepository(db)
	usersService := users.NewService(usersRepository)
	router = users.MakeUserRoutes(router, usersService)
	router.Mount("/api", router)

	level.Info(logger).Log("bootstrap", fmt.Sprintf("routes successfully initialized, now listening on port %d", parsedPost))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", parsedPost), router); err != nil {
		level.Error(logger).Log("bootstrap", "failed to start server", "err", err)
		os.Exit(1)
	}
}
