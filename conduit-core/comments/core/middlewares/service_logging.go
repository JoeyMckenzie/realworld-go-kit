package middlewares

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/core"
    commentsDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/comments"
    "time"
)

type commentsServiceLoggingMiddleware struct {
    logger log.Logger
    next   core.CommentsService
}

func NewCommentsServiceLoggingMiddleware(logger log.Logger) core.CommentsServiceMiddleware {
    return func(next core.CommentsService) core.CommentsService {
        return &commentsServiceLoggingMiddleware{
            logger: logger,
            next:   next,
        }
    }
}

func (mw *commentsServiceLoggingMiddleware) AddComment(ctx context.Context, request *commentsDomain.AddArticleCommentServiceRequest) (comment *commentsDomain.CommentDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "AddComment",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "AddComment",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.AddComment(ctx, request)
}

func (mw *commentsServiceLoggingMiddleware) DeleteComment(ctx context.Context, request *commentsDomain.DeleteArticleCommentServiceRequest) (err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "DeleteComment",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "DeleteComment",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.DeleteComment(ctx, request)
}

func (mw *commentsServiceLoggingMiddleware) GetArticleComments(ctx context.Context, request *commentsDomain.GetCommentsServiceRequest) (comments []*commentsDomain.CommentDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "GetArticleComments",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "GetArticleComments",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.GetArticleComments(ctx, request)
}
