package comments

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
    commentsDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/comments"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
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
            apiRequest := request.(commentsDomain.AddCommentApiRequest)

            serviceRequest := commentsDomain.AddArticleCommentServiceRequest{
                Body:   apiRequest.Comment.Body,
                UserId: tokenMeta.UserId,
                Slug:   apiRequest.Slug,
            }

            comment, err := service.AddComment(ctx, &serviceRequest)

            if err != nil {
                return nil, err
            }

            return &commentsDomain.CommentResponse{
                Comment: comment,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeDeleteCommentEndpoint(service core.CommentsService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            serviceRequest := request.(commentsDomain.DeleteArticleCommentServiceRequest)
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
        serviceRequest := request.(commentsDomain.GetCommentsServiceRequest)
        comments, err := service.GetArticleComments(ctx, &serviceRequest)

        if err != nil {
            return nil, err
        }

        return &commentsDomain.CommentsResponse{
            Comments: comments,
        }, nil
    }
}
