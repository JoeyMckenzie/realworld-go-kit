package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/internal/comments/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/comments/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/services"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
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
	userId, err := services.
		NewTokenService().
		GetRequiredUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	if err != nil {
		return nil, utilities.ErrUnauthorized
	}

	request := domain.AddArticleCommentServiceRequest{
		UserId: userId,
		Slug:   chi.URLParam(r, "slug"),
	}

	return request, nil
}

func decodeDeleteCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, err := services.
		NewTokenService().
		GetRequiredUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	if err != nil {
		return nil, utilities.ErrUnauthorized
	}

	commentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)

	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("comments", err)
	}

	request := domain.DeleteArticleCommentServiceRequest{
		UserId:    userId,
		Slug:      chi.URLParam(r, "slug"),
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
