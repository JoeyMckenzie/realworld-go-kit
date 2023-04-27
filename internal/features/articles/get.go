package articles

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
    "sync"
)

var (
    wg       = new(sync.WaitGroup)
    syncLock = new(sync.Mutex)
)

func (as *articlesService) GetFeed(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error) {
    articles, err := as.articlesRepository.GetArticles(ctx, userId, "", "", "", request.Limit, request.Offset)
    return as.getMappedArticles(ctx, articles, err)
}

func (as *articlesService) ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error) {
    articles, err := as.articlesRepository.GetArticles(ctx, userId, request.Tag, request.Author, request.Favorited, request.Limit, request.Offset)
    return as.getMappedArticles(ctx, articles, err)
}

func (as *articlesService) GetArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    article, err := as.articlesRepository.GetArticle(ctx, slug, userId)

    if err != nil && err != sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "error while attempting find article", "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        as.logger.ErrorCtx(ctx, "article not found", "slug", slug, "user_id", userId)
        return &domain.Article{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrArticlesNotFound)
    }

    tagList, err := as.tagsRepository.GetArticleTags(ctx, article.ID)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting retrieve tags for article", "slug", slug, "user_id", userId, "article_id", article.ID)
        return &domain.Article{}, nil
    }

    return article.ToArticle(tagList)
}

func (as *articlesService) getMappedArticles(ctx context.Context, articles []repositories.ArticleCompositeQuery, err error) ([]domain.Article, error) {
    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to retrieve articles", "err", err)
        return []domain.Article{}, err
    }

    if len(articles) == 0 {
        as.logger.InfoCtx(ctx, "no articles found")
        return []domain.Article{}, nil
    }

    var mappedArticles []domain.Article
    {
        tagQueryErrors := make(chan error, len(articles))

        // We'll capture errors that occur while querying for article tags in our go routines, propagating them back to the client
        for _, article := range articles {
            // Capture the current article to avoid unexpected derefs
            currentArticle := article
            wg.Add(1)

            // We'll wrap our call to add article tags in an anonymous go routine,
            // which we'll then communicate back any errors back to the error channel
            go func() {
                tagQueryErrors <- as.addArticleWithTags(ctx, currentArticle, &mappedArticles)
            }()
        }

        wg.Wait()

        if err = <-tagQueryErrors; err != nil {
            as.logger.ErrorCtx(ctx, "error while attempting to retrieve article tags", "errors", err)
            return []domain.Article{}, err
        }
    }

    return mappedArticles, nil
}

func (as *articlesService) addArticleWithTags(ctx context.Context, article repositories.ArticleCompositeQuery, mappedArticles *[]domain.Article) error {
    defer wg.Done()
    associatedTags, err := as.tagsRepository.GetArticleTags(ctx, article.ID)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to retrieve article tags", "err", err, "article_id", article.ID)
        return err
    }

    mappedArticle, err := article.ToArticle(associatedTags)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to map article", "err", err, "article_id", article.ID)
        return err
    }

    // Since we're mutating a shared list, only allow a single routine to update the list at a time
    syncLock.Lock()
    *mappedArticles = append(*mappedArticles, *mappedArticle)
    syncLock.Unlock()

    return nil
}
