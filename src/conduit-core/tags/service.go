package tags

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/tag"
)

type (
	TagsService interface {
		GetTags(ctx context.Context) ([]string, error)
	}

	tagsService struct {
		validator *validator.Validate
		client    *ent.Client
	}

	TagsServiceMiddleware func(tagsService TagsService) TagsService
)

func NewTagsService(validator *validator.Validate, client *ent.Client) TagsService {
	return &tagsService{
		validator: validator,
		client:    client,
	}
}

func (ts *tagsService) GetTags(ctx context.Context) ([]string, error) {
	tags, err := ts.client.Tag.
		Query().
		Order(ent.Desc(tag.FieldCreateTime)).
		Select(tag.FieldTag).
		Strings(ctx)

	if err != nil {
		return nil, shared.NewInternalServerErrorWithContext("tags", err)
	}

	return tags, nil
}
