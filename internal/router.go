package internal

import (
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/internal/features"
    articlesApi "github.com/joeymckenzie/realworld-go-kit/internal/features/articles"
    profilesApi "github.com/joeymckenzie/realworld-go-kit/internal/features/profiles"
    usersApi "github.com/joeymckenzie/realworld-go-kit/internal/features/users"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

// NewRouter initializes a new instance of the chi router, mounting all sub-routes for users, articles, etc.
func NewRouter(logger log.Logger, container *features.ServiceContainer) *chi.Mux {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    router.Use(middleware.RequestID)
    router.Use(shared.CorsPolicy)
    router.Use(shared.JsonContentType)
    router.Use(middleware.AllowContentType("application/json"))

    router = usersApi.MakeUserRoutes(logger, router, container.UsersService)
    router = profilesApi.MakeProfileRoutes(logger, router, container.ProfilesService)
    router = articlesApi.MakeArticlesRoutes(logger, router, container.ArticlesService)
    router.Mount("/api", router)

    return router
}
