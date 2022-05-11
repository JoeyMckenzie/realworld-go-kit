package conduit_api

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-kit/log"
    articlesApi "github.com/joeymckenzie/realworld-go-kit/conduit-api/articles"
    commentsApi "github.com/joeymckenzie/realworld-go-kit/conduit-api/comments"
    tagsApi "github.com/joeymckenzie/realworld-go-kit/conduit-api/tags"
    usersApi "github.com/joeymckenzie/realworld-go-kit/conduit-api/users"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewConduitRouter(logger log.Logger, serviceRegister *conduit_core.ConduitServiceRegister) *chi.Mux {
    router := api.NewChiRouter()
    router.Get("/metrics", promhttp.Handler().ServeHTTP)
    router = usersApi.MakeUsersTransport(router, logger, serviceRegister.UsersService)
    router = articlesApi.MakeArticlesTransport(router, logger, serviceRegister.ArticlesService)
    router = commentsApi.MakeCommentsTransport(router, logger, serviceRegister.CommentsService)
    router = tagsApi.MakeTagsTransport(router, logger, serviceRegister.TagsService)
    router.Mount("/api", router)

    return router
}
