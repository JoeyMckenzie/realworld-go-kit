package api

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/api"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

type commentEndpoints struct {
    MakeAddCommentEndpoint    endpoint.Endpoint
    MakeDeleteCommentEndpoint endpoint.Endpoint
    MakeGetCommentsEndpoint   endpoint.Endpoint
}

func NewCommentEndpoints(service core.CommentsService) *commentEndpoints {
    return &commentEndpoints{
        MakeAddCommentEndpoint:    makeAddCommentEndpoint(service),
        MakeDeleteCommentEndpoint: makeDeleteCommentEndpoint(service),
        MakeGetCommentsEndpoint:   makeGetCommentsEndpoint(service),
    }
}

func makeAddCommentEndpoint(service core.CommentsService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(domain.AddCommentApiRequest)

            serviceRequest := domain.AddArticleCommentServiceRequest{
                Body:   apiRequest.Comment.Body,
                UserId: tokenMeta.UserId,
                Slug:   apiRequest.Slug,
            }

            comment, err := service.AddComment(ctx, &serviceRequest)

            if err != nil {
                return nil, err
            }

            return &domain.CommentResponse{
                Comment: comment,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeDeleteCommentEndpoint(service core.CommentsService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            serviceRequest := request.(domain.DeleteArticleCommentServiceRequest)
            serviceRequest.UserId = tokenMeta.UserId

            if err := service.DeleteComment(ctx, &serviceRequest); err != nil {
                return nil, err
            }

            return nil, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeGetCommentsEndpoint(service core.CommentsService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        serviceRequest := request.(domain.GetCommentsServiceRequest)
        comments, err := service.GetArticleComments(ctx, &serviceRequest)

        if err != nil {
            return nil, err
        }

        return &domain.CommentsResponse{
            Comments: comments,
        }, nil
    }
}
