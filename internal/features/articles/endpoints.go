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
