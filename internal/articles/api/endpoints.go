package api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func makeCreateArticlesEndpoint(service core.ArticlesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if tokenMeta, ok := ctx.Value(api.TokenMeta{}).(api.TokenMeta); ok && tokenMeta.UserId > 0 {
			apiRequest := request.(domain.UpsertArticleApiRequest)
			response, err := service.CreateArticle(ctx, &domain.UpsertArticleServiceRequest{
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
