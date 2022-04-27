package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
	"strconv"
)

func MakeArticlesTransport(router *chi.Mux, logger log.Logger, service core.ArticlesService) *chi.Mux {
	endpoints := NewArticleEndpoints(service)

	createArticleHandler := httpTransport.NewServer(
		endpoints.MakeCreateArticleEndpoint,
		decodeCreateArticleRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	getArticlesHandler := httpTransport.NewServer(
		endpoints.MakeGetArticlesEndpoint,
		decodeGetArticlesRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	getArticleHandler := httpTransport.NewServer(
		endpoints.MakeGetArticleEndpoint,
		decodeGetArticleRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	getFeedHandler := httpTransport.NewServer(
		endpoints.MakeGetFeedEndpoint,
		decodeGetFeedRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	updateArticleHandler := httpTransport.NewServer(
		endpoints.MakeUpdateArticleEndpoint,
		decodeUpdateArticleRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	deleteArticleHandler := httpTransport.NewServer(
		endpoints.MakeDeleteArticleEndpoint,
		decodeDeleteArticleRequest,
		api.EncodeSuccessfulResponseWithNoContent,
		api.HandlerOptions(logger)...,
	)

	favoriteArticleHandler := httpTransport.NewServer(
		endpoints.MakeFavoriteArticleEndpoint,
		decodeFavoriteArticleRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	unfavoriteArticleHandler := httpTransport.NewServer(
		endpoints.MakeUnfavoriteArticleEndpoint,
		decodeFavoriteArticleRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	router.Route("/articles", func(r chi.Router) {
		r.Get("/", getArticlesHandler.ServeHTTP)

		r.Route("/{slug}", func(r chi.Router) {
			r.Get("/", getArticleHandler.ServeHTTP)
			r.Group(func(r chi.Router) {
				r.Use(api.AuthorizedRequestMiddleware)
				r.Put("/", updateArticleHandler.ServeHTTP)
				r.Delete("/", deleteArticleHandler.ServeHTTP)
				r.Post("/favorite", favoriteArticleHandler.ServeHTTP)
				r.Delete("/favorite", unfavoriteArticleHandler.ServeHTTP)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(api.AuthorizedRequestMiddleware)
			r.Post("/", createArticleHandler.ServeHTTP)
			r.Get("/feed", getFeedHandler.ServeHTTP)
		})
	})

	return router
}

func decodeCreateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.CreateArticleApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	return request, nil
}

func decodeUpdateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.UpdateArticleApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	request.Article.Slug = chi.URLParam(r, "slug")

	return request, nil
}

func decodeGetArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, _ := services.
		NewTokenService().
		GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	request := domain.GetArticleServiceRequest{
		UserId: userId,
		Slug:   chi.URLParam(r, "slug"),
	}

	return request, nil
}

func decodeGetArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	limit, limitOk := getDefaultParamValue(20, r.URL.Query().Get("limit"))
	offset, offsetOk := getDefaultParamValue(0, r.URL.Query().Get("offset"))

	if !limitOk || !offsetOk {
		return nil, utilities.ErrInvalidLimitOrOffsetValue
	}

	userId, _ := services.
		NewTokenService().
		GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	request := domain.GetArticlesServiceRequest{
		UserId:    userId,
		Tag:       r.URL.Query().Get("tag"),
		Author:    r.URL.Query().Get("author"),
		Favorited: r.URL.Query().Get("favorited"),
		Limit:     limit,
		Offset:    offset,
	}

	return request, nil
}

func decodeGetFeedRequest(_ context.Context, r *http.Request) (interface{}, error) {
	limit, limitOk := getDefaultParamValue(20, r.URL.Query().Get("limit"))
	offset, offsetOk := getDefaultParamValue(0, r.URL.Query().Get("offset"))

	if !limitOk || !offsetOk {
		return nil, utilities.ErrInvalidLimitOrOffsetValue
	}

	userId, err := services.
		NewTokenService().
		GetRequiredUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	if err != nil {
		return nil, utilities.ErrUnauthorized
	}

	request := domain.GetArticlesServiceRequest{
		UserId: userId,
		Limit:  limit,
		Offset: offset,
	}
	return request, nil
}

func decodeDeleteArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, _ := services.
		NewTokenService().
		GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	request := domain.DeleteArticleServiceRequest{
		UserId:      userId,
		ArticleSlug: chi.URLParam(r, "slug"),
	}

	return request, nil
}

func decodeFavoriteArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, err := services.
		NewTokenService().
		GetRequiredUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	if err != nil {
		return nil, utilities.ErrUnauthorized
	}

	request := domain.ArticleFavoriteServiceRequest{
		UserId: userId,
		Slug:   chi.URLParam(r, "slug"),
	}

	return request, nil
}

func getDefaultParamValue(defaultValue int, queryParamValue string) (int, bool) {
	value := defaultValue

	if queryParamValue != "" {
		parsedQueryParamValue, err := strconv.ParseInt(queryParamValue, 10, 64)
		if err != nil {
			return value, false
		}

		value = int(parsedQueryParamValue)
	}

	return value, true
}
