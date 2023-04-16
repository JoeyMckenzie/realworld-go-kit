package articles

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/data"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type (
    ArticlesService interface {
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
    }

    articlesService struct {
        logger             log.Logger
        articlesRepository data.ArticlesRepository
        usersRepository    data.UsersRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(logger log.Logger, articlesRepository data.ArticlesRepository, usersRepository data.UsersRepository) ArticlesService {
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
    if err != nil {
        return &domain.Article{}, shared.MakeApiError(err)
    }

    // Next, start a transaction, so we can rollback anytime one of the inserts fails
    tx, err := as.articlesRepository.BeginTransaction(ctx)
    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while attempting to begin transaction", "article_title", request.Title, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    createdTags := make([]uuid.UUID, len(request.TagList))

    // next, we'll roll through all the tags on the request and create new tags that don't exist
    for _, tag := range request.TagList {
        createdTag, err := as.articlesRepository.CreateTag(ctx, tx, tag)

        if err != nil {
            level.Error(as.logger).Log(loggingSpan, "error while adding tag", "tag", tag)
            return &domain.Article{}, shared.MakeApiError(tx.Rollback())
        }

        createdTags = append(createdTags, createdTag.ID)
    }

    // Next, we'll create the article
    createdArticle, err := as.articlesRepository.CreateArticle(ctx, tx, request.Title, request.Title, request.Description, request.Body)

    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while creating article", "slug", request.Title)
        return &domain.Article{}, shared.MakeApiError(tx.Rollback())
    }

    // Finally, create the associated article tags
    for _, tagId := range createdTags {
        if _, err := as.articlesRepository.CreateArticleTag(ctx, tx, tagId, createdArticle.ID); err != nil {
            level.Error(as.logger).Log(loggingSpan, "error while creating article tag", "article_id", createdArticle.ID, "tag_id", tagId)
            return &domain.Article{}, shared.MakeApiError(tx.Rollback())
        }
    }

    return createdArticle.ToArticle(author.ToUser(""), request.TagList, 0, false, false), nil
}
