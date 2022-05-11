package tags

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/tags"
    tagsDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/tags"
)

type tagEndpoints struct {
    MakeGetTagsEndpoint endpoint.Endpoint
}

func NewTagEndpoints(service tags.TagsService) *tagEndpoints {
    return &tagEndpoints{
        MakeGetTagsEndpoint: makeGetTagsEndpoint(service),
    }
}

func makeGetTagsEndpoint(service tags.TagsService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        tagsResponse, err := service.GetTags(ctx)

        if err != nil {
            return nil, err
        }

        return &tagsDomain.GetTagsResponse{
            Tags: tagsResponse,
        }, nil
    }
}
