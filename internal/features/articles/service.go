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

    // First, grab the associated user
    author, err := as.usersRepository.GetUserById(ctx, authorId)

    if err != nil && err != sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "error while attempting to retrieve author", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "no author found while attempting to create article", "author_id", authorId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

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
        createdTag, err := as.articlesRepository.CreateTag(ctx, tx, tag)

        if err != nil {
            level.Error(as.logger).Log(loggingSpan, "error while adding tag", "tag", tag)
            return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
        }

        createdTags = append(createdTags, createdTag.ID)
    }

    // Then, verify the slug is unique - if not, add a unique hash to the slug
    articleSlug := slug.Make(articleRequest.Title)
    existingArticle, err := as.articlesRepository.GetArticleBySlug(ctx, tx, articleSlug)

    if err != nil && err != sql.ErrNoRows {
        level.Error(as.logger).Log(loggingSpan, "error while verifying slug uniqueness", "slug", articleSlug, "err", err)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    } else if existingArticle.Slug != "" {
        articleSlug = fmt.Sprintf("%s-%s", articleSlug, uuid.New().String())
    }

    // Next, we'll create the article
    createdArticle, err := as.articlesRepository.CreateArticle(ctx, tx, authorId, articleSlug, articleRequest.Title, articleRequest.Description, articleRequest.Body)

    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while creating article", "slug", articleRequest.Title)
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    // Finally, create the associated article tags and commit the transaction
    for _, tagId := range createdTags {
        if _, err := as.articlesRepository.CreateArticleTag(ctx, tx, tagId, createdArticle.ID); err != nil {
            level.Error(as.logger).Log(loggingSpan, "error while creating article tag", "article_id", createdArticle.ID, "tag_id", tagId)
            return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
        }
    }

    if err = tx.Commit(); err != nil {
        return &domain.Article{}, shared.MakeApiErrorWithFallback(err, tx.Rollback())
    }

    return createdArticle.ToArticle(author.ToProfile(false), request.Article.TagList, 0, false, false)
}
