package main

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	articlesApi "github.com/joeymckenzie/realworld-go-kit/internal/articles/api"
	articlesCore "github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	articlesMiddlewares "github.com/joeymckenzie/realworld-go-kit/internal/articles/core/middlewares"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersApi "github.com/joeymckenzie/realworld-go-kit/internal/users/api"
	usersCore "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	usersMiddlewares "github.com/joeymckenzie/realworld-go-kit/internal/users/core/middlewares"
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

	var usersService usersCore.UsersService
	{
		requestCount, requestLatency := utilities.NewServiceMetrics("users_service")
		usersService = usersCore.NewUsersService(requestValidator, usersRepository, services.NewTokenService(), services.NewSecurityService())
		usersService = usersMiddlewares.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = usersMiddlewares.NewUsersServiceMetrics(requestCount, requestLatency)(usersService)
		usersService = usersMiddlewares.NewUsersServiceRequestValidationMiddleware(logger, requestValidator)(usersService)
	}

	var articlesService articlesCore.ArticlesService
	{
		requestCount, requestLatency := utilities.NewServiceMetrics("articles_service")
		articlesService = articlesCore.NewArticlesServices(requestValidator, articlesRepository, usersRepository)
		articlesService = articlesMiddlewares.NewArticlesServiceLoggingMiddleware(logger)(articlesService)
		articlesService = articlesMiddlewares.NewArticlesServiceMetrics(requestCount, requestLatency)(articlesService)
		articlesService = articlesMiddlewares.NewArticlesServiceRequestValidationMiddleware(logger, requestValidator)(articlesService)
	}

	router := api.NewChiRouter()
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router = usersApi.MakeUsersTransport(router, logger, usersService)
	router = articlesApi.MakeArticlesTransport(router, logger, articlesService)
	router.Mount("/api", router)

	if err = http.ListenAndServe(":8080", router); err != nil {
		level.Error(logger).Log("main", "error during server startup", "error", err)
		os.Exit(1)
	}
}
