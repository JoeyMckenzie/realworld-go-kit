package middlewares

import (
    "context"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/domain"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

type commentsServiceRequestValidationMiddleware struct {
    logger    log.Logger
    validator *validator.Validate
    next      core.CommentsService
}

func NewCommentsServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) core.CommentsServiceMiddleware {
    return func(next core.CommentsService) core.CommentsService {
        return &commentsServiceRequestValidationMiddleware{
            logger:    logger,
            validator: validator,
            next:      next,
        }
    }
}

func (mw *commentsServiceRequestValidationMiddleware) AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (*domain.CommentDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.AddComment(ctx, request)
}

func (mw *commentsServiceRequestValidationMiddleware) DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) error {
    if err := mw.validator.Struct(request); err != nil {
        return api.NewValidationError(err)
    }

    return mw.next.DeleteComment(ctx, request)
}

func (mw *commentsServiceRequestValidationMiddleware) GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) ([]*domain.CommentDto, error) {
    if request == nil {
        return nil, utilities.ErrNilInput
    }

    return mw.next.GetArticleComments(ctx, request)
}
