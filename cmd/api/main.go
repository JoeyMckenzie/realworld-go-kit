package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersApi "github.com/joeymckenzie/realworld-go-kit/internal/users/api"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core/middlewares"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

const postgresDiver = "postgres"

func main() {
	// Build out our logging instance
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"caller", log.DefaultCaller,
		)
		level.Info(logger).Log("main", "bootstrapping application...")
		defer level.Info(logger).Log("main", "application shutting down...")
	}

	// Load in configuration
	if err := godotenv.Load(); err != nil {
		level.Error(logger).Log(
			"configuration", "error while bootstrapping environment",
			"error", err,
		)
		os.Exit(1)
	}

	// Build out our connection to the database
	db, err := sqlx.Open(postgresDiver, os.Getenv("CONNECTION_STRING"))

	if err != nil {
		level.Error(logger).Log(
			"database", "error while connecting to postgres",
			"error", err,
		)
		os.Exit(1)
	}

	requestValidator := validator.New()

	var usersRepository usersPersistence.UsersRepository
	{
		usersRepository = usersPersistence.NewUsersRepository(db)
		usersRepository = usersPersistence.NewUsersRepositoryLoggingMiddleware(logger)(usersRepository)
	}

	var articlesRepository articlesPersistence.ArticlesRepository
	{
		articlesRepository = articlesPersistence.NewArticlesRepository(db)
		articlesRepository = articlesPersistence.NewArticlesRepositoryLoggingMiddleware(logger)(articlesRepository)
	}

	var usersService core.UsersService
	{
		requestCount, requestLatency := utilities.NewServiceMetrics("users_service")
		usersService = core.NewUsersService(requestValidator, usersRepository, services.NewTokenService(), services.NewSecurityService())
		usersService = middlewares.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = middlewares.NewUsersServiceMetrics(requestCount, requestLatency)(usersService)
		usersService = middlewares.NewUsersServiceRequestValidationMiddleware(logger, requestValidator)(usersService)
	}

	router := api.MakeChiRouter()
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router.Mount("/api", makeApiTransports(logger, usersService, requestValidator))

	if err = http.ListenAndServe(":8080", router); err != nil {
		level.Error(logger).Log("main", "error during server startup", "error", err)
		os.Exit(1)
	}
}

func makeApiTransports(logger log.Logger, usersService core.UsersService, validator *validator.Validate) http.Handler {
	router := chi.NewRouter()
	router.Mount("/", usersApi.MakeUsersTransport(logger, usersService, validator))
	return router
}
