package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	articlesApi "github.com/joeymckenzie/realworld-go-kit/internal/articles/api"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	usersApi "github.com/joeymckenzie/realworld-go-kit/internal/users/api"
	usersCore "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitializeRouter(logger log.Logger, usersService usersCore.UsersService, articlesService core.ArticlesService) *chi.Mux {
	router := api.NewChiRouter()
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router = usersApi.MakeUsersTransport(router, logger, usersService)
	router = articlesApi.MakeArticlesTransport(router, logger, articlesService)
	router.Mount("/api", router)

	return router
}
