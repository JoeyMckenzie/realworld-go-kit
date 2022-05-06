package api

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/pkg/api"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

type articleEndpoints struct {
    MakeCreateArticleEndpoint     endpoint.Endpoint
    MakeGetArticlesEndpoint       endpoint.Endpoint
    MakeGetArticleEndpoint        endpoint.Endpoint
    MakeGetFeedEndpoint           endpoint.Endpoint
    MakeUpdateArticleEndpoint     endpoint.Endpoint
    MakeDeleteArticleEndpoint     endpoint.Endpoint
    MakeFavoriteArticleEndpoint   endpoint.Endpoint
    MakeUnfavoriteArticleEndpoint endpoint.Endpoint
}

func NewArticleEndpoints(service core.ArticlesService) *articleEndpoints {
    return &articleEndpoints{
        MakeCreateArticleEndpoint:     makeCreateArticleEndpoint(service),
        MakeGetArticlesEndpoint:       makeGetArticlesEndpoint(service),
        MakeGetArticleEndpoint:        makeGetArticleEndpoint(service),
        MakeGetFeedEndpoint:           makeGetFeedEndpoint(service),
        MakeUpdateArticleEndpoint:     makeUpdateArticleEndpoint(service),
        MakeDeleteArticleEndpoint:     makeDeleteArticleEndpoint(service),
        MakeFavoriteArticleEndpoint:   makeFavoriteArticleEndpoint(service),
        MakeUnfavoriteArticleEndpoint: makeUnfavoriteArticleEndpoint(service),
    }
}

func makeCreateArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(domain.CreateArticleApiRequest)
            response, err := service.CreateArticle(ctx, &domain.CreateArticleServiceRequest{
                UserId:      tokenMeta.UserId,
                Title:       apiRequest.Article.Title,
                Description: apiRequest.Article.Description,
                Body:        apiRequest.Article.Body,
                TagList:     apiRequest.Article.TagList,
            })

            if err != nil {
                return nil, err
            }

            return &domain.UpsertArticleResponse{
                Article: response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeGetArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.GetArticleServiceRequest)
        article, err := service.GetArticle(ctx, &apiRequest)

        if err != nil {
            return nil, err
        }

        return &domain.GetArticleResponse{
            Article: article,
        }, nil
    }
}

func makeGetArticlesEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.GetArticlesServiceRequest)
        articles, err := service.GetArticles(ctx, &apiRequest)

        if err != nil {
            return nil, err
        }

        return &domain.GetArticlesResponse{
            Articles:      articles,
            ArticlesCount: len(articles),
        }, nil
    }
}

func makeGetFeedEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        apiRequest := request.(domain.GetArticlesServiceRequest)
        articles, err := service.GetFeed(ctx, &apiRequest)

        if err != nil {
            return nil, err
        }

        return &domain.GetArticlesResponse{
            Articles:      articles,
            ArticlesCount: len(articles),
        }, nil
    }
}

func makeUpdateArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(domain.UpdateArticleApiRequest)
            response, err := service.UpdateArticle(ctx, &domain.UpdateArticleServiceRequest{
                UserId:      tokenMeta.UserId,
                ArticleSlug: apiRequest.Article.Slug,
                Title:       apiRequest.Article.Title,
                Description: apiRequest.Article.Description,
                Body:        apiRequest.Article.Body,
            })

            if err != nil {
                return nil, err
            }

            return &domain.UpsertArticleResponse{
                Article: response,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeDeleteArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            apiRequest := request.(domain.DeleteArticleServiceRequest)
            err := service.DeleteArticle(ctx, &apiRequest)

            if err != nil {
                return nil, err
            }

            return nil, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeFavoriteArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            serviceRequest := request.(domain.ArticleFavoriteServiceRequest)
            article, err := service.FavoriteArticle(ctx, &serviceRequest)

            if err != nil {
                return nil, err
            }

            return &domain.GetArticleResponse{
                Article: article,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}

func makeUnfavoriteArticleEndpoint(service core.ArticlesService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
            serviceRequest := request.(domain.ArticleFavoriteServiceRequest)
            article, err := service.UnfavoriteArticle(ctx, &serviceRequest)

            if err != nil {
                return nil, err
            }

            return &domain.GetArticleResponse{
                Article: article,
            }, nil
        }

        return nil, utilities.ErrUnauthorized
    }
}
