package articles

import (
    "context"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "sync"
)

var (
    wg       = new(sync.WaitGroup)
    syncLock = new(sync.Mutex)
)

func (as *articlesService) ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error) {
    articles, err := as.articlesRepository.GetArticles(ctx, userId, request.Tag, request.Author, request.Favorited, request.Limit, request.Offset)

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
    associatedTags, err := as.articlesRepository.GetArticleTags(ctx, article.ID)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to retrieve article tags", "err", err, "article_id", article.ID)
        return err
    }

    mappedArticle, err := article.ToArticle(associatedTags)

    if err != nil {
        as.logger.ErrorCtx(ctx, "error while attempting to map article", "err", err, "article_id", article.ID)
        return err
    }

    syncLock.Lock()
    *mappedArticles = append(*mappedArticles, *mappedArticle)
    syncLock.Unlock()

    return nil
}
