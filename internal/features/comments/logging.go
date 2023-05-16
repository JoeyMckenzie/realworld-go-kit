package comments

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"golang.org/x/exp/slog"
)

type commentsServiceLoggingMiddleware struct {
	logger *slog.Logger
	next   CommentsService
}

func NewCommentsServiceLoggingMiddleware(logger *slog.Logger) CommentsServiceMiddleware {
	return func(next CommentsService) CommentsService {
		return &commentsServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw commentsServiceLoggingMiddleware) AddComment(ctx context.Context, request domain.CreateCommentRequest, userId uuid.UUID) (comment *domain.Comment, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"AddComment",
			"request_time", time.Since(begin),
			"error", err,
			"comment_added", comment != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"AddComment",
		"slug", request.Slug,
		"user_id", userId,
	)

	return mw.next.AddComment(ctx, request, userId)
}

func (mw commentsServiceLoggingMiddleware) DeleteComment(ctx context.Context, userId, commentId uuid.UUID) (err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"DeleteComment",
			"request_time", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"DeleteComment",
		"user_id", userId,
		"comment_id", commentId,
	)

	return mw.next.DeleteComment(ctx, userId, commentId)
}

func (mw commentsServiceLoggingMiddleware) GetComments(ctx context.Context, request domain.CommentRetrievalRequest, userId uuid.UUID) (comments []domain.Comment, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"GetComments",
			"request_time", time.Since(begin),
			"error", err,
			"comment_found", len(comments),
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"GetComments",
		"user_id", userId,
	)

	return mw.next.GetComments(ctx, request, userId)
}
