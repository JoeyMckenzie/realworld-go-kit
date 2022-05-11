package tags

import (
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/tags"
)

func MakeTagsTransport(router *chi.Mux, logger log.Logger, service tags.TagsService) *chi.Mux {
	endpoints := NewTagEndpoints(service)

	getTagsHandler := httpTransport.NewServer(
		endpoints.MakeGetTagsEndpoint,
		apiUtilities.DecodeDefaultRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	router.Get("/tags", getTagsHandler.ServeHTTP)

	return router
}
