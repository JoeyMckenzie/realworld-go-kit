package articles

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/google/uuid"
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (as *articlesService) UpdateArticle(ctx context.Context, request domain.UpdateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    // First, verify the article exists
    article, err := as.articlesRepository.GetArticle(ctx, request.Slug, authorId)

    if err != nil && err != sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "error while attempting to update article", "slug", request.Slug, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "article not found while attempting to update", "slug", request.Slug, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    // Next, verify the author of the article is the user on the request
    if article.UserID != authorId {
        as.logger.ErrorCtx(ctx, "user attempting to update non-authored article", "slug", request.Slug, "author_id", article.UserID, "user_id_on_request", authorId)
        // Rather than return a 403, it's often times better to obfuscate non-owned resource mutation attempts
        // with a not found response to help combat attacks from malicious users looking for probing for resources
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    // Next, update the applicable fields on the request
    updatedTitle := shared.GetUpdatedValueIfApplicable(request.Article.Title, article.Title)
    updatedBody := shared.GetUpdatedValueIfApplicable(request.Article.Body, article.Body)
    updatedDescription := shared.GetUpdatedValueIfApplicable(request.Article.Description, article.Description)

    // Slugify the title, and update it if we find a clashing slug
    updatedSlug := slug.Make(updatedTitle)
    existingArticleWithClashingSlug, err := as.articlesRepository.GetArticle(ctx, updatedSlug, authorId)

    if err != nil && err != sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "error while attempting to check for clashing article slug", "slug", updatedSlug, "article_id", article.ID)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if existingArticleWithClashingSlug.Slug != "" {
        updatedSlug = slug.Make(fmt.Sprintf("%s-%s", updatedSlug, uuid.New().String()))
    }

    updatedArticle, err := as.articlesRepository.UpdateArticle(ctx, article.ID, authorId, updatedSlug, updatedTitle, updatedBody, updatedDescription)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while updating article", "slug", updatedSlug, "article_id", article.ID)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    // Finally, grab the associated tags
    associatedTags, err := as.tagsRepository.GetArticleTags(ctx, updatedArticle.ID)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while updating article", "slug", updatedSlug, "article_id", article.ID)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    return updatedArticle.ToArticle(associatedTags)
}
