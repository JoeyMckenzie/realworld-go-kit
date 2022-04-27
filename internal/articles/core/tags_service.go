package core

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
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
		Select(tag.FieldTag).
		Strings(ctx)

	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("tags", err)
	}

	return tags, nil
}
