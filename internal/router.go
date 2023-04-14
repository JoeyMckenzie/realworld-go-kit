package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	profilesApi "github.com/joeymckenzie/realworld-go-kit/internal/profiles/api"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	usersApi "github.com/joeymckenzie/realworld-go-kit/internal/users/api"
)

// NewRouter initializes a new instance of the chi router, mounting all sub-routes for users, articles, etc.
func NewRouter(logger log.Logger, container *ServiceContainer) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(shared.CorsPolicy)
	router.Use(shared.JsonContentType)
	router.Use(middleware.AllowContentType("application/json"))

	router = usersApi.MakeUserRoutes(logger, router, container.UsersService)
	router = profilesApi.MakeProfileRoutes(logger, router, container.ProfilesService)
	router.Mount("/api", router)

	return router
}
