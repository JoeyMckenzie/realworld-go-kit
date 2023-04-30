package articles

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (as *articlesService) DeleteArticle(ctx context.Context, slug string, authorId uuid.UUID) error {
    // First, verify the article exists
    article, err := as.articlesRepository.GetArticle(ctx, nil, slug, authorId)

    if shared.IsValidSqlErr(err) {
        as.logger.ErrorCtx(ctx, "error while attempting to delete article", "slug", slug, "author_id", authorId)
        return shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "article not found while attempting to delete", "slug", slug, "author_id", authorId)
        return shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    // Next, verify the author of the article is the user on the request
    if article.UserID != authorId {
        as.logger.ErrorCtx(ctx, "user attempting to update non-authored article", "slug", slug, "author_id", article.UserID, "user_id_on_request", authorId)
        // Rather than return a 403, it's often times better to obfuscate non-owned resource mutation attempts
        // with a not found response to help combat attacks from malicious users looking for probing for resources
        return shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    return as.articlesRepository.DeleteArticle(ctx, article.ID)
}
