package conduit_api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/conduit-api/articles"
	"github.com/joeymckenzie/realworld-go-kit/conduit-api/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-api/tags"
	"github.com/joeymckenzie/realworld-go-kit/conduit-api/users"
	apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
	conduitCore "github.com/joeymckenzie/realworld-go-kit/conduit-core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewConduitRouter(logger log.Logger, serviceRegister *conduitCore.ConduitServiceRegister) *chi.Mux {
	router := apiUtilities.NewChiRouter()
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router = users.MakeUsersTransport(router, logger, serviceRegister.UsersService)
	router = articles.MakeArticlesTransport(router, logger, serviceRegister.ArticlesService)
	router = comments.MakeCommentsTransport(router, logger, serviceRegister.CommentsService)
	router = tags.MakeTagsTransport(router, logger, serviceRegister.TagsService)
	router.Mount("/api", router)

	return router
}
