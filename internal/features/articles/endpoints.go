package articles

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
)

func makeCreateArticleEndpoint(service ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        uuidClaim := ctx.Value(utilities.TokenContextKey{}).(utilities.TokenContextKey)
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
