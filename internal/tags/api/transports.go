package api

import (
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/tags/core"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
)

func MakeTagsTransport(router *chi.Mux, logger log.Logger, service core.TagsService) *chi.Mux {
	endpoints := NewTagEndpoints(service)

	getTagsHandler := httpTransport.NewServer(
		endpoints.MakeGetTagsEndpoint,
		api.DecodeDefaultRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	router.Get("/tags", getTagsHandler.ServeHTTP)

	return router
}