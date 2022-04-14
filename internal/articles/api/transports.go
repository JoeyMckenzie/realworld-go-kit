package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

func MakeArticlesTransport(router *chi.Mux, logger log.Logger, service core.ArticlesService) *chi.Mux {
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(api.EncodeError),
	}

	endpoints := NewArticleEndpoints(service)

	createArticleHandler := httpTransport.NewServer(
		endpoints.MakeCreateArticleEndpoint,
		decodeUpsertArticleRequest,
		api.EncodeSuccessfulResponse,
		options...,
	)

	getArticleHandler := httpTransport.NewServer(
		endpoints.MakeGetArticlesEndpoint,
		decodeGetArticlesRequest,
		api.EncodeSuccessfulResponse,
		options...,
	)

	router.Route("/articles", func(r chi.Router) {
		r.Get("/", getArticleHandler.ServeHTTP)
		r.Group(func(innerRouter chi.Router) {
			innerRouter.Use(api.AuthorizedRequestMiddleware)
			innerRouter.Post("/", createArticleHandler.ServeHTTP)
		})
	})

	return router
}

func decodeUpsertArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.UpsertArticleApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	return request, nil
}

func decodeGetArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := domain.GetArticlesServiceRequest{
        Tag:       "",
        Author:    "",
        Favorited: "",
        Limit:     20,
        Offset:    0,
    }
	return request, nil
}
