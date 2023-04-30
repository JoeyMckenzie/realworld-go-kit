package articles

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (as *articlesService) FavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    return as.handleFavoriteRequest(ctx, true, slug, userId)
}

func (as *articlesService) UnavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    return as.handleFavoriteRequest(ctx, false, slug, userId)
}

func (as *articlesService) handleFavoriteRequest(ctx context.Context, favorite bool, slug string, userId uuid.UUID) (*domain.Article, error) {
    existingArticle, err := as.articlesRepository.GetArticle(ctx, nil, slug, userId)

    if shared.IsValidSqlErr(err) {
        as.logger.ErrorCtx(ctx, "error while attempting to verify article exists for favorite", "favorite", favorite, "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "no article found to favorite", "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    if favorite {
        err = as.articlesRepository.AddFavorite(ctx, existingArticle.ID, userId)
    } else {
        err = as.articlesRepository.DeleteFavorite(ctx, existingArticle.ID, userId)
    }

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to favorite article", "favorite", favorite, "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    // Re-retrieve the article and tags
    existingArticle, err := as.articlesRepository.GetArticle(ctx, nil, slug, userId)
    articleTags, err := as.tagsRepository.GetArticleTags(ctx, existingArticle.ID)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to retrieve article tags while favoriting", "favorite", favorite, "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    return existingArticle.ToArticle(articleTags)
}
