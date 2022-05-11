package articles

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/articles"
	articlesDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"net/http"
	"strconv"
)

func MakeArticlesTransport(router *chi.Mux, logger log.Logger, service articles.ArticlesService) *chi.Mux {
	endpoints := NewArticleEndpoints(service)

	createArticleHandler := httpTransport.NewServer(
		endpoints.MakeCreateArticleEndpoint,
		decodeCreateArticleRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	getArticlesHandler := httpTransport.NewServer(
		endpoints.MakeGetArticlesEndpoint,
		decodeGetArticlesRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	getArticleHandler := httpTransport.NewServer(
		endpoints.MakeGetArticleEndpoint,
		decodeGetArticleRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	getFeedHandler := httpTransport.NewServer(
		endpoints.MakeGetFeedEndpoint,
		decodeGetFeedRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	updateArticleHandler := httpTransport.NewServer(
		endpoints.MakeUpdateArticleEndpoint,
		decodeUpdateArticleRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	deleteArticleHandler := httpTransport.NewServer(
		endpoints.MakeDeleteArticleEndpoint,
		decodeDeleteArticleRequest,
		apiUtilities.EncodeSuccessfulResponseWithNoContent,
		apiUtilities.HandlerOptions(logger)...,
	)

	favoriteArticleHandler := httpTransport.NewServer(
		endpoints.MakeFavoriteArticleEndpoint,
		decodeFavoriteArticleRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	unfavoriteArticleHandler := httpTransport.NewServer(
		endpoints.MakeUnfavoriteArticleEndpoint,
		decodeFavoriteArticleRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	router.Route("/articles", func(r chi.Router) {
		r.Get("/", getArticlesHandler.ServeHTTP)

		r.Route("/{slug}", func(r chi.Router) {
			r.Get("/", getArticleHandler.ServeHTTP)
			r.Group(func(r chi.Router) {
				r.Use(apiUtilities.AuthorizedRequestMiddleware)
				r.Put("/", updateArticleHandler.ServeHTTP)
				r.Delete("/", deleteArticleHandler.ServeHTTP)
				r.Post("/favorite", favoriteArticleHandler.ServeHTTP)
				r.Delete("/favorite", unfavoriteArticleHandler.ServeHTTP)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(apiUtilities.AuthorizedRequestMiddleware)
			r.Post("/", createArticleHandler.ServeHTTP)
			r.Get("/feed", getFeedHandler.ServeHTTP)
		})
	})

	return router
}

func decodeCreateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request articlesDomain.CreateArticleApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	return request, nil
}

func decodeUpdateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request articlesDomain.UpdateArticleApiRequest

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

	request := articlesDomain.GetArticleServiceRequest{
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

	request := articlesDomain.GetArticlesServiceRequest{
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

	request := articlesDomain.GetArticlesServiceRequest{
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

	request := articlesDomain.DeleteArticleServiceRequest{
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

	request := articlesDomain.ArticleFavoriteServiceRequest{
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
