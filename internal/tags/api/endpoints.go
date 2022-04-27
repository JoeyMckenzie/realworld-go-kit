package api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/joeymckenzie/realworld-go-kit/internal/tags/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/tags/domain"
)

type tagEndpoints struct {
	MakeGetTagsEndpoint endpoint.Endpoint
}

func NewTagEndpoints(service core.TagsService) *tagEndpoints {
	return &tagEndpoints{
		MakeGetTagsEndpoint: makeGetTagsEndpoint(service),
	}
}

func makeGetTagsEndpoint(service core.TagsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		tags, err := service.GetTags(ctx)

		if err != nil {
			return nil, err
		}

		return &domain.GetTagsResponse{
			Tags: tags,
		}, nil
	}
}
