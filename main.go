package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/app/users"
	_ "github.com/lib/pq"
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

	if err != nil {
		level.Error(logger).Log("main", "failed to parse port", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("main", "initializing database connection...")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		level.Error(logger).Log("main", "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		level.Error(logger).Log("main", "failed to ping database", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("main", "database connection successfully initialized, building routes")

	userService := users.NewService(db)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router = users.MakeEndpoints(router, userService)

	level.Info(logger).Log("main", fmt.Sprintf("routes successfully initialized, now listening on port %d", parsedPost))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", parsedPost), router); err != nil {
		level.Error(logger).Log("main", "failed to start server", "err", err)
		os.Exit(1)
	}
}
