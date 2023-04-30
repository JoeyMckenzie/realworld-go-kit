package articles

import (
    "context"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "golang.org/x/exp/slog"
)

type (
    ArticlesService interface {
        GetArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error)
        GetFeed(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error)
        ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error)
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
        UpdateArticle(ctx context.Context, request domain.UpdateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
        DeleteArticle(ctx context.Context, slug string, authorId uuid.UUID) error
        FavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error)
        UnavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error)
        GetArticleTags(ctx context.Context) ([]string, error)
    }

    articlesService struct {
        logger             *slog.Logger
        articlesRepository repositories.ArticlesRepository
        usersRepository    repositories.UsersRepository
        tagsRepository     repositories.TagsRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(
    logger *slog.Logger,
    articlesRepository repositories.ArticlesRepository,
    usersRepository repositories.UsersRepository,
    tagsRepository repositories.TagsRepository) ArticlesService {
    return &articlesService{
        logger:             logger,
        articlesRepository: articlesRepository,
        usersRepository:    usersRepository,
        tagsRepository:     tagsRepository,
    }
}

func (as *articlesService) GetArticleTags(ctx context.Context) ([]string, error) {
    tags, err := as.tagsRepository.GetTags(ctx)

    if shared.IsValidSqlErr(err) {
        as.logger.ErrorCtx(ctx, "error while querying for tags")
        return []string{}, shared.MakeApiError(err)
    }

    return tags, nil
}
