package core

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

type (
	ArticlesService interface {
		CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (*domain.ArticleDto, error)
	}

	articlesServices struct {
		validator  *validator.Validate
		repository persistence.ArticlesRepository
	}

	ArticlesServiceMiddleware func(articlesService ArticlesService) ArticlesService
)

func NewArticlesServices(validator *validator.Validate, repository persistence.ArticlesRepository) ArticlesService {
	return &articlesServices{
		validator:  validator,
		repository: repository,
	}
}

func (as *articlesServices) CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (*domain.ArticleDto, error) {
	// Verify the article title slug is unique
	articleSlug := slug.Make(request.Title)
	existingArticleFromSlug, err := as.repository.FindArticleBySlug(ctx, articleSlug)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if existingArticleFromSlug != nil {
		return nil, api.NewApiErrorWithContext(http.StatusConflict, "article", utilities.ErrArticleTitleExists)
	}

	// Create the article
	createdArticle, err := as.repository.CreateArticle(ctx, request.UserId, request.Title, articleSlug, request.Description, request.Body)
	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("article", err)
	}

	var articleTagIdsToCreate []int
	var returnedTagList []string

	// Create any tags on the request
	if request.TagList != nil && len(*request.TagList) > 0 {
		returnedTagList = *request.TagList

		for _, tag := range *request.TagList {
			// Check for an existing tag
			existingTag, err := as.repository.GetTag(ctx, tag)

			if err != nil && err != sql.ErrNoRows {
				return nil, api.NewInternalServerErrorWithContext("tags", err)
			}

			// If an existing tag is not found, created it and add it to the list of associated article tags to create
			if existingTag == nil {
				existingTag, err = as.repository.CreateTag(ctx, tag)
				if err != nil {
					return nil, api.NewInternalServerErrorWithContext("tags", err)
				}
			}

			articleTagIdsToCreate = append(articleTagIdsToCreate, existingTag.Id)
		}
	}

	// Finally, create the tags associated with the article
	if len(articleTagIdsToCreate) > 0 {
		for _, tagId := range articleTagIdsToCreate {
			if _, err = as.repository.CreateArticleTag(ctx, tagId, createdArticle.Id); err != nil {
				return nil, api.NewInternalServerErrorWithContext("articleTags", err)
			}
		}
	}

	return &domain.ArticleDto{
		Slug:           articleSlug,
		Title:          createdArticle.Title,
		Description:    createdArticle.Description,
		Body:           createdArticle.Body,
		TagList:        returnedTagList,
		CreatedAt:      createdArticle.CreatedAt,
		UpdatedAt:      createdArticle.UpdatedAt,
		Favorited:      false,
		FavoritesCount: 0,
		Author:         domain.AuthorDto{},
	}, nil
}
