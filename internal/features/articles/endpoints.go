package articles

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

func makeCreateArticleEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        createArticleRequest := request.(domain.CreateArticleRequest)
        createdArticle, err := service.CreateArticle(ctx, createArticleRequest, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ArticleResponse{
            Article: createdArticle,
        }, nil
    }
}

func makeUpdateArticleEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        updateArticleRequest := request.(domain.UpdateArticleRequest)
        updatedArticle, err := service.UpdateArticle(ctx, updateArticleRequest, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ArticleResponse{
            Article: updatedArticle,
        }, nil
    }
}

func makeDeleteArticleEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        articleRequest := request.(domain.ArticleRetrievalRequest)
        return nil, service.DeleteArticle(ctx, articleRequest.Slug, uuidClaim.UserId)
    }
}

func makeListArticlesEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        listArticlesRequest := request.(domain.ListArticlesRequest)
        articles, err := service.ListArticles(ctx, listArticlesRequest, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ArticlesResponse{
            Articles:      articles,
            ArticlesCount: len(articles),
        }, nil
    }
}

func makeFeedArticlesEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        listArticlesRequest := request.(domain.ListArticlesRequest)
        articles, err := service.GetFeed(ctx, listArticlesRequest, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ArticlesResponse{
            Articles:      articles,
            ArticlesCount: len(articles),
        }, nil
    }
}

func makeGetArticleEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
        articleRequest := request.(domain.ArticleRetrievalRequest)
        article, err := service.GetArticle(ctx, articleRequest.Slug, uuidClaim.UserId)

        if err != nil {
            return nil, err
        }

        return &domain.ArticleResponse{
            Article: article,
        }, nil
    }
}
