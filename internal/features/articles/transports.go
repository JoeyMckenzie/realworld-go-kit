package articles

import (
    "context"
    "encoding/json"
    "github.com/go-chi/chi"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "golang.org/x/exp/slog"
    "net/http"
    "strconv"
)

const (
    defaultLimit  = 20
    defaultOffset = 0
)

func MakeArticlesRoutes(logger *slog.Logger, router *chi.Mux, service ArticlesService) *chi.Mux {
    createArticleHandler := httptransport.NewServer(
        makeCreateArticleEndpoint(service),
        decodeCreateArticleRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    listArticleHandler := httptransport.NewServer(
        makeListArticlesEndpoint(service),
        decodeListArticlesRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    router.Route("/articles", func(r chi.Router) {
        r.Group(func(r chi.Router) {
            r.Use(shared.AuthorizationOptional)
            r.Get("/", listArticleHandler.ServeHTTP)
        })

        r.Group(func(r chi.Router) {
            r.Use(shared.AuthorizationRequired)
            r.Post("/", createArticleHandler.ServeHTTP)
        })
    })

    return router
}

func decodeCreateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request domain.CreateArticleRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, shared.ErrorWithContext("error while attempting to decode the article request", shared.ErrInvalidRequestBody)
    }

    return request, nil
}

func decodeListArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
    limit := defaultLimit
    offset := defaultOffset
    tag := r.URL.Query().Get("tag")
    author := r.URL.Query().Get("author")
    favorited := r.URL.Query().Get("favorited")
    requestLimit := r.URL.Query().Get("limit")
    requestOffset := r.URL.Query().Get("offset")

    var err error
    {
        if requestLimit != "" {
            limit, err = strconv.Atoi(requestLimit)

            // If an error occurs, reset to the default
            if err != nil {
                limit = defaultLimit
            }
        }

        if requestOffset != "" {
            offset, err = strconv.Atoi(requestLimit)

            // If an error occurs, reset to the default
            if err != nil {
                offset = defaultOffset
            }
        }
    }

    request := domain.ListArticlesRequest{
        Limit:     limit,
        Offset:    offset,
        Tag:       tag,
        Author:    author,
        Favorited: favorited,
    }

    return request, nil
}
