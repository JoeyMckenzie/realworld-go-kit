package conduit_core

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-kit/log"
    articlesApi "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/api"
    commentsApi "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/api"
    tagsApi "github.com/joeymckenzie/realworld-go-kit/conduit-core/tags/api"
    usersApi "github.com/joeymckenzie/realworld-go-kit/conduit-core/users/api"
    "github.com/joeymckenzie/realworld-go-kit/pkg/api"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitializeRouter(logger log.Logger, serviceRegister *ConduitServiceRegister) *chi.Mux {
    router := api.NewChiRouter()
    router.Get("/metrics", promhttp.Handler().ServeHTTP)
    router = usersApi.MakeUsersTransport(router, logger, serviceRegister.usersService)
    router = articlesApi.MakeArticlesTransport(router, logger, serviceRegister.articlesService)
    router = commentsApi.MakeCommentsTransport(router, logger, serviceRegister.commentsService)
    router = tagsApi.MakeTagsTransport(router, logger, serviceRegister.tagsService)
    router.Mount("/api", router)

    return router
}
