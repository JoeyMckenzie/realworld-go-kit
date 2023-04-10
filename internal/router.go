package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
)

// NewRouter initializes a new instance of the chi router, mounting all sub-routes for users, articles, etc.
func NewRouter(logger log.Logger, container *ServiceContainer) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router = users.MakeUserRoutes(router, container.UsersService)

	router.Mount("/api", router)

	return router
}
