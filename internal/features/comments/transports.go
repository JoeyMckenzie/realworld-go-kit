package comments

import (
    "context"
    "encoding/json"
    "github.com/go-chi/chi"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gofrs/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "golang.org/x/exp/slog"
    "net/http"
)

func MakeCommentsRoutes(logger *slog.Logger, router *chi.Mux, service CommentsService) *chi.Mux {
    addCommentHandler := httptransport.NewServer(
        makeAddCommentEndpoint(service),
        decodeAddCommentRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    deleteCommentHandler := httptransport.NewServer(
        makeDeleteCommentEndpoint(service),
        decodeDeleteCommentRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    getCommentsHandler := httptransport.NewServer(
        makeGetArticleCommentsEndpoint(service),
        decodeGetArticleCommentsRequest,
        shared.EncodeSuccessfulOkResponse,
        shared.HandlerOptions(logger)...,
    )

    router.Route("/articles/{slug}/comments", func(r chi.Router) {
        r.Group(func(r chi.Router) {
            r.Use(shared.AuthorizationOptional)
            r.Get("/", getCommentsHandler.ServeHTTP)
        })

        r.Group(func(r chi.Router) {
            r.Use(shared.AuthorizationRequired)
            r.Post("/", addCommentHandler.ServeHTTP)
            r.Delete("/{id}", deleteCommentHandler.ServeHTTP)
        })
    })

    return router
}

func decodeAddCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request domain.CreateCommentRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, shared.ErrorWithContext("error while attempting to decode the comment request", shared.ErrInvalidRequestBody)
    }

    request.Slug = chi.URLParam(r, "slug")

    return request, nil
}

func decodeDeleteCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
    requestId := chi.URLParam(r, "id")
    parsedId, err := uuid.FromString(requestId)

    if err != nil {
        return nil, shared.ErrorWithContext("error while attempting to parse the comment ID", err)
    }

    return domain.CommentRetrievalRequest{
        Slug: chi.URLParam(r, "slug"),
        ID:   parsedId,
    }, nil
}

func decodeGetArticleCommentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return domain.CommentRetrievalRequest{
        Slug: chi.URLParam(r, "slug"),
    }, nil
}
