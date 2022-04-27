package internal

import (
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	articlesCore "github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	articlesMiddlewares "github.com/joeymckenzie/realworld-go-kit/internal/articles/core/middlewares"
	usersCore "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	usersMiddlewares "github.com/joeymckenzie/realworld-go-kit/internal/users/core/middlewares"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func InitializeServices(logger log.Logger, entClient *ent.Client) (usersCore.UsersService, articlesCore.ArticlesService) {
	// Build out services
	requestValidator := validator.New()

	var usersService usersCore.UsersService
	{
		requestCount, requestLatency := utilities.NewServiceMetrics("users_service")
		usersService = usersCore.NewUsersService(requestValidator, entClient, services.NewTokenService(), services.NewSecurityService())
		usersService = usersMiddlewares.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = usersMiddlewares.NewUsersServiceMetrics(requestCount, requestLatency)(usersService)
		usersService = usersMiddlewares.NewUsersServiceRequestValidationMiddleware(logger, requestValidator)(usersService)
	}

	var articlesService articlesCore.ArticlesService
	{
		requestCount, requestLatency := utilities.NewServiceMetrics("articles_service")
		articlesService = articlesCore.NewArticlesServices(requestValidator, entClient)
		articlesService = articlesMiddlewares.NewArticlesServiceLoggingMiddleware(logger)(articlesService)
		articlesService = articlesMiddlewares.NewArticlesServiceMetrics(requestCount, requestLatency)(articlesService)
		articlesService = articlesMiddlewares.NewArticlesServiceRequestValidationMiddleware(logger, requestValidator)(articlesService)
	}

	return usersService, articlesService
}
