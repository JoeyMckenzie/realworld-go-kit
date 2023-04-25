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
)

func MakeArticlesRoutes(logger *slog.Logger, router *chi.Mux, service ArticlesService) *chi.Mux {
    createArticleHandler := httptransport.NewServer(
        makeCreateArticleEndpoint(service),
        decodeCreateArticleRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    router.Route("/articles", func(r chi.Router) {
        r.Use(shared.AuthorizationRequired)
        r.Post("/", createArticleHandler.ServeHTTP)
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
