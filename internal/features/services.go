package features

import (
    "github.com/go-playground/validator/v10"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/features/articles"
    "github.com/joeymckenzie/realworld-go-kit/internal/features/profiles"
    "github.com/joeymckenzie/realworld-go-kit/internal/features/users"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
    "golang.org/x/exp/slog"
)

type ServiceContainer struct {
    UsersService    users.UsersService
    ProfilesService profiles.ProfilesService
    ArticlesService articles.ArticlesService
}

// NewServiceContainer builds the downstream services used throughout the application.
func NewServiceContainer(logger *slog.Logger, db *sqlx.DB) *ServiceContainer {
    validation := validator.New()
    usersRepository := repositories.NewUsersRepository(db)

    var usersService users.UsersService
    {
        tokenService := utilities.NewTokenService()
        securityService := utilities.NewSecurityService()
        usersService = users.NewUsersService(logger, usersRepository, tokenService, securityService)
        usersService = users.NewUsersServiceLoggingMiddleware(logger)(usersService)
        usersService = users.NewUsersServiceValidationMiddleware(validation)(usersService)
    }

    var profilesService profiles.ProfilesService
    {
        profilesService = profiles.NewProfileService(logger, usersRepository)
        profilesService = profiles.NewProfileServiceLoggingMiddleware(logger)(profilesService)
    }

    var articlesService articles.ArticlesService
    {
        articlesRepository := repositories.NewArticlesRepository(db)
        articlesService = articles.NewArticlesService(logger, articlesRepository, usersRepository)
        articlesService = articles.NewArticlesServiceValidationMiddleware(validation)(articlesService)
    }

    return &ServiceContainer{
        UsersService:    usersService,
        ProfilesService: profilesService,
        ArticlesService: articlesService,
    }
}
