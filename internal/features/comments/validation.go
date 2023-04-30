package comments

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type commentsServiceValidationMiddleware struct {
    validator *validator.Validate
    next      CommentsService
}

func NewCommentsServiceValidationMiddleware(validator *validator.Validate) CommentsServiceMiddleware {
    return func(next CommentsService) CommentsService {
        return &commentsServiceValidationMiddleware{
            validator: validator,
            next:      next,
        }
    }
}

func (mw commentsServiceValidationMiddleware) AddComment(ctx context.Context, request domain.CreateCommentRequest, userId uuid.UUID) (*domain.Comment, error) {
    if err := mw.validator.StructCtx(ctx, request); err != nil {
        return &domain.Comment{}, shared.MakeValidationError(err)
    }

    return mw.next.AddComment(ctx, request, userId)
}

func (mw commentsServiceValidationMiddleware) DeleteComment(ctx context.Context, userId, commentId uuid.UUID) error {
    return mw.next.DeleteComment(ctx, userId, commentId)
}

func (mw commentsServiceValidationMiddleware) GetComments(ctx context.Context, request domain.CommentRetrievalRequest, userId uuid.UUID) ([]domain.Comment, error) {
    if err := mw.validator.StructCtx(ctx, request); err != nil {
        return []domain.Comment{}, shared.MakeValidationError(err)
    }

    return mw.next.GetComments(ctx, request, userId)
}
