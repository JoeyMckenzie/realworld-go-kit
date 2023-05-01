package comments

import (
    "context"
    "database/sql"
    "github.com/gofrs/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (cs *commentsService) GetComments(ctx context.Context, request domain.CommentRetrievalRequest, userId uuid.UUID) ([]domain.Comment, error) {
    // First, verify the article exists
    if _, err := cs.articlesRepository.GetArticle(ctx, nil, request.Slug, userId); shared.IsValidSqlErr(err) {
        cs.logger.ErrorCtx(ctx, "error while retrieving article for comments", "slug", request.Slug, "user_id", userId)
        return []domain.Comment{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        cs.logger.WarnCtx(ctx, "article not found for comments", "slug", request.Slug, "user_id", userId)
        return []domain.Comment{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    // Next, grab all the comments using the optional user on the request to determine following status
    articleComments, err := cs.commentsRepository.GetArticleComments(ctx, userId, request.Slug)

    if err != nil {
        cs.logger.ErrorCtx(ctx, "error while attempting to retrieve article comments", "user_id", userId, "slug", request.Slug)
        return []domain.Comment{}, shared.MakeApiError(err)
    }

    // Finally, map the comment query models into our response model
    var comments []domain.Comment
    {
        for _, articleComment := range articleComments {
            mappedComment, err := articleComment.ToComment()

            if err != nil {
                cs.logger.WarnCtx(ctx, "error while attempting to map comment, skipping", "comment_id", articleComment.ID, "slug", request.Slug)
                continue
            }

            comments = append(comments, *mappedComment)
        }
    }

    return comments, nil
}
