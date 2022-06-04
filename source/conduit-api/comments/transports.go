package comments

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	apiUtilities "github.com/joeymckenzie/realworld-go-kit/conduit-api/utilities"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	commentsDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/services"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"net/http"
	"strconv"
)

func MakeCommentsTransport(router *chi.Mux, logger log.Logger, service comments.CommentsService) *chi.Mux {
	endpoints := NewCommentEndpoints(service)

	addCommentHandler := httpTransport.NewServer(
		endpoints.MakeAddCommentEndpoint,
		decodeAddCommentRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	deleteCommentHandler := httpTransport.NewServer(
		endpoints.MakeDeleteCommentEndpoint,
		decodeDeleteCommentRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	getCommentsHandler := httpTransport.NewServer(
		endpoints.MakeGetCommentsEndpoint,
		decodeGetCommentsRequest,
		apiUtilities.EncodeSuccessfulResponse,
		apiUtilities.HandlerOptions(logger)...,
	)

	router.Route("/articles/{slug}/comments", func(r chi.Router) {
		r.Get("/", getCommentsHandler.ServeHTTP)
		r.Group(func(r chi.Router) {
			r.Use(apiUtilities.AuthorizedRequestMiddleware)
			r.Post("/", addCommentHandler.ServeHTTP)
			r.Delete("/{id:^[1-9]+}", deleteCommentHandler.ServeHTTP)
		})
	})

	return router
}

func decodeAddCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request commentsDomain.AddCommentApiRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utilities.ErrInvalidRequestBody
	}

	request.Slug = chi.URLParam(r, "slug")

	return request, nil
}

func decodeDeleteCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	commentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)

	if err != nil {
		return nil, shared.NewInternalServerErrorWithContext("commentsDomain", err)
	}

	request := commentsDomain.DeleteArticleCommentServiceRequest{
		CommentId: int(commentId),
	}

	return request, nil
}

func decodeGetCommentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userId, _ := services.
		NewTokenService().
		GetOptionalUserIdFromAuthorizationHeader(r.Header.Get("Authorization"))

	request := commentsDomain.GetCommentsServiceRequest{
		UserId: userId,
		Slug:   chi.URLParam(r, "slug"),
	}

	return request, nil
}
