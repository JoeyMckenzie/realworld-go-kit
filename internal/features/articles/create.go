package articles

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/gofrs/uuid"
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (as *articlesService) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    as.logger.InfoCtx(ctx, "verifying user exists before creating article", "author_id", authorId)

    // First, check that the author exists in the database
    _, err := as.usersRepository.GetUserById(ctx, authorId)

    if shared.IsValidSqlErr(err) {
        as.logger.ErrorCtx(ctx, "error while attempting to retrieve author", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "no author found while attempting to create article", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    as.logger.InfoCtx(ctx, "user successfully identified", "author_id", authorId)

    // Next, start a transaction, so we can rollback anytime one of the inserts fails
    tx, err := as.articlesRepository.BeginTransaction(ctx)
    articleRequest := request.Article

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to begin transaction", "article_title", articleRequest.Title, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    // Keep a reference to all the created tags, if any, to add to the article tags table
    var createdTags []uuid.UUID
    {
        // Next, we'll roll through all the tags on the request and create new tags that don't exist
        for _, tag := range articleRequest.TagList {
            as.logger.InfoCtx(ctx, "attempting to create tag", "tag", tag, "author_id", authorId, "article_title", articleRequest.Title)
            createdTag, err := as.tagsRepository.CreateTag(ctx, tx, tag)

            if err != nil {
                as.logger.ErrorCtx(ctx, "error while adding tag", "tag", tag, "err", err)
                return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
            }

            as.logger.InfoCtx(ctx, "tag successfully created", "tag", tag, "author_id", authorId, "article_title", articleRequest.Title)
            createdTags = append(createdTags, createdTag.ID)
        }
    }

    as.logger.InfoCtx(ctx, "verifying slug uniqueness", "author_id", authorId, "article_title", articleRequest.Title)

    // Then, verify the slug is unique - if not, add a unique hash to the slug
    articleSlug := slug.Make(articleRequest.Title)
    existingArticle, err := as.articlesRepository.GetArticle(ctx, tx, articleSlug, authorId)

    if shared.IsValidSqlErr(err) {
        as.logger.ErrorCtx(ctx, "error while verifying slug uniqueness", "slug", articleSlug, "err", err)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    } else if existingArticle != nil && existingArticle.Slug != "" {
        as.logger.InfoCtx(ctx, "slug is not unique, overriding with unique ID", "author_id", authorId, "article_title", articleRequest.Title, "slug", articleSlug)
        articleSlug = fmt.Sprintf("%s-%s", articleSlug, uuid.Must(uuid.NewV4()))
    }

    // Next, we'll create the article
    as.logger.InfoCtx(ctx, "attempting to create article", "author_id", authorId, "article_title", articleRequest.Title, "slug", articleSlug)
    createdArticle, err := as.articlesRepository.CreateArticle(ctx, tx, authorId, articleSlug, articleRequest.Title, articleRequest.Description, articleRequest.Body)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while creating article", "slug", articleRequest.Title)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    as.logger.InfoCtx(ctx, "article successfully created", "author_id", authorId, "article_id", createdArticle.ID, "article_title", createdArticle.Title, "slug", createdArticle.Slug)

    // Finally, create the associated article tags and commit the transaction
    for _, tagId := range createdTags {
        if tagId != uuid.Nil {
            as.logger.InfoCtx(ctx, "attempting to add article tags", "author_id", authorId, "article_id", createdArticle.ID, "tag_id", tagId)

            if _, err = as.tagsRepository.CreateArticleTag(ctx, tx, tagId, createdArticle.ID); err != nil {
                as.logger.ErrorCtx(ctx, "error while creating article tag", "article_id", createdArticle.ID, "tag_id", tagId)
                return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
            }

            as.logger.InfoCtx(ctx, "tag successfully added", "author_id", authorId, "article_id", createdArticle.ID, "tag_id", tagId)
        }
    }

    as.logger.InfoCtx(ctx, "all process successfully completed, saving transaction", "author_id", authorId, "article_id", createdArticle.ID)

    if err = tx.Commit(); err != nil {
        as.logger.InfoCtx(ctx, "error while attempting to save transaction", "author_id", authorId, "article_id", createdArticle.ID, "err", err)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    as.logger.InfoCtx(ctx, "transaction successfully saved", "author_id", authorId, "article_id", createdArticle.ID)

    return createdArticle.ToArticle(request.Article.TagList)
}
