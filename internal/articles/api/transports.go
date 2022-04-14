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
    "strconv"
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
        r.Group(func(r chi.Router) {
            r.Use(api.AuthorizedRequestMiddleware)
            r.Post("/", createArticleHandler.ServeHTTP)
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
    var limit int
    {
        limit = 20

        if limitQueryParam := r.URL.Query().Get("limit"); limitQueryParam != "" {
            parsedLimit, err := strconv.ParseInt(limitQueryParam, 10, 64)

            if err != nil {
                return nil, utilities.ErrInvalidLimitOrOffsetValue
            }

            limit = int(parsedLimit)
        }
    }

    var offset int
    {
        offset = 0

        if offsetQueryParam := r.URL.Query().Get("offset"); offsetQueryParam != "" {
            parsedOffset, err := strconv.ParseInt(offsetQueryParam, 10, 64)

            if err != nil {
                return nil, utilities.ErrInvalidLimitOrOffsetValue
            }

            offset = int(parsedOffset)
        }
    }

    request := domain.GetArticlesServiceRequest{
        Tag:       r.URL.Query().Get("tag"),
        Author:    r.URL.Query().Get("author"),
        Favorited: r.URL.Query().Get("favorited"),
        Limit:     limit,
        Offset:    offset,
    }
    return request, nil
}
