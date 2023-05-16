package comments

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func makeAddCommentEndpoint(service CommentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
		commentRequest := request.(domain.CreateCommentRequest)
		comment, err := service.AddComment(ctx, commentRequest, uuidClaim.UserId)
		if err != nil {
			return nil, err
		}

		return &domain.CommentResponse{
			Comment: comment,
		}, nil
	}
}

func makeGetArticleCommentsEndpoint(service CommentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
		commentRequest := request.(domain.CommentRetrievalRequest)
		comments, err := service.GetComments(ctx, commentRequest, uuidClaim.UserId)
		if err != nil {
			return nil, err
		}

		return &domain.CommentsResponse{
			Comments: comments,
		}, nil
	}
}

func makeDeleteCommentEndpoint(service CommentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
		commentRequest := request.(domain.CommentRetrievalRequest)

		if err := service.DeleteComment(ctx, uuidClaim.UserId, commentRequest.ID); err != nil {
			return nil, err
		}

		return nil, nil
	}
}
