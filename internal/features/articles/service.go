package articles

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

type (
    ArticlesService interface {
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
    }

    articlesService struct {
        logger             log.Logger
        articlesRepository repositories.ArticlesRepository
        usersRepository    repositories.UsersRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(logger log.Logger, articlesRepository repositories.ArticlesRepository, usersRepository repositories.UsersRepository) ArticlesService {
    return &articlesService{
        logger:             logger,
        articlesRepository: articlesRepository,
        usersRepository:    usersRepository,
    }
}

func (as *articlesService) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    const loggingSpan string = "create_article"

    level.Info(as.logger).Log(loggingSpan, "verifying user exists before creating article", "author_id", authorId)

    // First, grab the associated user
    author, err := as.usersRepository.GetUserById(ctx, authorId)

    if err != nil && err != sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "error while attempting to retrieve author", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "no author found while attempting to create article", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    level.Info(as.logger).Log(loggingSpan, "user successfully identified", "author_id", authorId)

    // Next, start a transaction, so we can rollback anytime one of the inserts fails
    tx, err := as.articlesRepository.BeginTransaction(ctx)
    articleRequest := request.Article

    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while attempting to begin transaction", "article_title", articleRequest.Title, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    // Keep a reference to all the created tags, if any, to add to the article tags table
    createdTags := make([]uuid.UUID, len(articleRequest.TagList))

    // Next, we'll roll through all the tags on the request and create new tags that don't exist
    for _, tag := range articleRequest.TagList {
        level.Info(as.logger).Log(loggingSpan, "attempting to create tag", "tag", tag, "author_id", authorId, "article_title", articleRequest.Title)
        createdTag, err := as.articlesRepository.CreateTag(ctx, tx, tag)

        if err != nil {
            level.Error(as.logger).Log(loggingSpan, "error while adding tag", "tag", tag, "err", err)
            return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
        }

        level.Info(as.logger).Log(loggingSpan, "tag successfully created", "tag", tag, "author_id", authorId, "article_title", articleRequest.Title)
        createdTags = append(createdTags, createdTag.ID)
    }

    level.Info(as.logger).Log(loggingSpan, "verifying slug uniqueness", "author_id", authorId, "article_title", articleRequest.Title)

    // Then, verify the slug is unique - if not, add a unique hash to the slug
    articleSlug := slug.Make(articleRequest.Title)
    existingArticle, err := as.articlesRepository.GetArticleBySlug(ctx, tx, articleSlug)

    if err != nil && err != sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "error while verifying slug uniqueness", "slug", articleSlug, "err", err)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    } else if existingArticle.Slug != "" {
        level.Info(as.logger).Log(loggingSpan, "slug is not unique, overriding with unique ID", "author_id", authorId, "article_title", articleRequest.Title, "slug", articleSlug)
        articleSlug = fmt.Sprintf("%s-%s", articleSlug, uuid.New().String())
    }

    // Next, we'll create the article
    level.Info(as.logger).Log(loggingSpan, "attempting to create article", "author_id", authorId, "article_title", articleRequest.Title, "slug", articleSlug)
    createdArticle, err := as.articlesRepository.CreateArticle(ctx, tx, authorId, articleSlug, articleRequest.Title, articleRequest.Description, articleRequest.Body)

    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while creating article", "slug", articleRequest.Title)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    level.Info(as.logger).Log(loggingSpan, "article successfully created", "author_id", authorId, "article_id", createdArticle.ID, "article_title", createdArticle.Title, "slug", createdArticle.Slug)

    // Finally, create the associated article tags and commit the transaction
    for _, tagId := range createdTags {
        if tagId != uuid.Nil {
            level.Info(as.logger).Log(loggingSpan, "attempting to add article tags", "author_id", authorId, "article_id", createdArticle.ID, "tag_id", tagId)

            if _, err = as.articlesRepository.CreateArticleTag(ctx, tx, tagId, createdArticle.ID); err != nil {
                level.Error(as.logger).Log(loggingSpan, "error while creating article tag", "article_id", createdArticle.ID, "tag_id", tagId)
                return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
            }

            level.Info(as.logger).Log(loggingSpan, "tag successfully added", "author_id", authorId, "article_id", createdArticle.ID, "tag_id", tagId)
        }
    }

    level.Info(as.logger).Log(loggingSpan, "all process successfully completed, saving transaction", "author_id", authorId, "article_id", createdArticle.ID)

    if err = tx.Commit(); err != nil {
        level.Info(as.logger).Log(loggingSpan, "error while attempting to save transaction", "author_id", authorId, "article_id", createdArticle.ID, "err", err)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    level.Info(as.logger).Log(loggingSpan, "transaction successfully saved", "author_id", authorId, "article_id", createdArticle.ID)

    return createdArticle.ToArticle(author.ToProfile(false), request.Article.TagList, 0, false, false)
}
