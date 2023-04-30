package comments

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (cs *commentsService) AddComment(ctx context.Context, request domain.CreateCommentRequest, userId uuid.UUID) (*domain.Comment, error) {
    // First, verify the article exists
    existingArticle, err := cs.articlesRepository.GetArticle(ctx, nil, request.Slug, userId)

    if shared.IsValidSqlErr(err) {
        cs.logger.ErrorCtx(ctx, "error while attempting to find existing article to add comment", "slug", request.Slug, "user_id", userId)
        return &domain.Comment{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        cs.logger.ErrorCtx(ctx, "no article found to add comment", "slug", request.Slug, "user_id", userId)
        return &domain.Comment{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    // Create the comment, map it back to the expected response
    newComment, err := cs.commentsRepository.AddComment(ctx, existingArticle.ID, userId, request.Comment.Body)

    if err != nil {
        cs.logger.ErrorCtx(ctx, "error while attempting to add comment to article", "slug", request.Slug, "user_id", userId, "article_id", existingArticle.ID)
        return &domain.Comment{}, shared.MakeApiError(err)
    }

    return newComment.ToComment()
}
