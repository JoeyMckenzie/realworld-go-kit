package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
)

// NewRouter initializes a new instance of the chi router, mounting all sub-routes for users, articles, etc.
func NewRouter(logger log.Logger, db *pgxpool.Pool, validation *validator.Validate) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	var usersService users.UsersService
	{
		usersRepository := users.NewRepository(db)
		usersService = users.NewService(usersRepository)
		usersService = users.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = users.NewUsersServiceValidationMiddleware(validation)(usersService)
	}

	router = users.MakeUserRoutes(router, usersService)

	router.Mount("/api", router)

	return router
}
