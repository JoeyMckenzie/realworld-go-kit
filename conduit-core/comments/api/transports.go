package api

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/domain"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"net/http"
	"strconv"
)

func MakeCommentsTransport(router *chi.Mux, logger log.Logger, service core.CommentsService) *chi.Mux {
	endpoints := NewCommentEndpoints(service)

	addCommentHandler := httpTransport.NewServer(
		endpoints.MakeAddCommentEndpoint,
		decodeAddCommentRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	deleteCommentHandler := httpTransport.NewServer(
		endpoints.MakeDeleteCommentEndpoint,
		decodeDeleteCommentRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	getCommentsHandler := httpTransport.NewServer(
		endpoints.MakeGetCommentsEndpoint,
		decodeGetCommentsRequest,
		api.EncodeSuccessfulResponse,
		api.HandlerOptions(logger)...,
	)

	router.Route("/articles/{slug}/comments", func(r chi.Router) {
		r.Get("/", getCommentsHandler.ServeHTTP)
		r.Group(func(r chi.Router) {
			r.Use(api.AuthorizedRequestMiddleware)
			r.Post("/", addCommentHandler.ServeHTTP)
			r.Delete("/{id:^[1-9]+}", deleteCommentHandler.ServeHTTP)
		})
	})

	return router
}

func decodeAddCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.AddCommentApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	request.Slug = chi.URLParam(r, "slug")

	return request, nil
}

func decodeDeleteCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	commentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)

	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("comments", err)
	}

	request := domain.DeleteArticleCommentServiceRequest{
		CommentId: int(commentId),
	}

	return request, nil
}

func decodeGetCommentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, _ := services.
		NewTokenService().
		GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	request := domain.GetCommentsServiceRequest{
		UserId: userId,
		Slug:   chi.URLParam(r, "slug"),
	}

	return request, nil
}
