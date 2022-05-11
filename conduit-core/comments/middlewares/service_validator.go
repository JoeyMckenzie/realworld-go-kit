package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	commentsDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

type commentsServiceRequestValidationMiddleware struct {
	logger    log.Logger
	validator *validator.Validate
	next      comments.CommentsService
}

func NewCommentsServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) comments.CommentsServiceMiddleware {
	return func(next comments.CommentsService) comments.CommentsService {
		return &commentsServiceRequestValidationMiddleware{
			logger:    logger,
			validator: validator,
			next:      next,
		}
	}
}

func (mw *commentsServiceRequestValidationMiddleware) AddComment(ctx context.Context, request *commentsDomain.AddArticleCommentServiceRequest) (*commentsDomain.CommentDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.AddComment(ctx, request)
}

func (mw *commentsServiceRequestValidationMiddleware) DeleteComment(ctx context.Context, request *commentsDomain.DeleteArticleCommentServiceRequest) error {
	if err := mw.validator.Struct(request); err != nil {
		return shared.NewValidationError(err)
	}

	return mw.next.DeleteComment(ctx, request)
}

func (mw *commentsServiceRequestValidationMiddleware) GetArticleComments(ctx context.Context, request *commentsDomain.GetCommentsServiceRequest) ([]*commentsDomain.CommentDto, error) {
	if request == nil {
		return nil, utilities.ErrNilInput
	}

	return mw.next.GetArticleComments(ctx, request)
}
