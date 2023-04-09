package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
)

// NewRouter initializes a new instance of the chi router, mounting all subroutes for users, articles, etc.
func NewRouter(db *pgxpool.Pool) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	usersRepository := users.NewRepository(db)
	usersService := users.NewService(usersRepository)
	router = users.MakeUserRoutes(router, usersService)

	router.Mount("/api", router)

	return router
}
