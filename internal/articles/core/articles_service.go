package core

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

type (
	ArticlesService interface {
		GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) ([]*domain.ArticleDto, error)
		CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (*domain.ArticleDto, error)
	}

	articlesServices struct {
		validator       *validator.Validate
		repository      persistence.ArticlesRepository
		usersRepository usersPersistence.UsersRepository
	}

	ArticlesServiceMiddleware func(articlesService ArticlesService) ArticlesService
)

func NewArticlesServices(validator *validator.Validate, repository persistence.ArticlesRepository, usersRepository usersPersistence.UsersRepository) ArticlesService {
	return &articlesServices{
		validator:       validator,
		repository:      repository,
		usersRepository: usersRepository,
	}
}

func (as *articlesServices) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) ([]*domain.ArticleDto, error) {
	return nil, nil
}

func (as *articlesServices) CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (*domain.ArticleDto, error) {
	// Verify the user exists, ensure no article is created without a valid existing user
	existingUser, err := as.usersRepository.GetUser(ctx, request.UserId)
	if err != nil || existingUser == nil {
		return nil, api.NewApiErrorWithContext(http.StatusConflict, "articles", utilities.ErrUserNotFound)
	}

	// Verify the article title slug is unique
	articleSlug := slug.Make(request.Title)
	existingArticleFromSlug, err := as.repository.FindArticleBySlug(ctx, articleSlug)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if existingArticleFromSlug != nil {
		return nil, api.NewApiErrorWithContext(http.StatusConflict, "article", utilities.ErrArticleTitleExists)
	}

	// Dedupe the request list of tags, initializing an empty list if none is provided on the request
	tagsToCreate := removeDuplicates(request.TagList)

	var articleTagsToCreate []int
	{
		// Create any tags on the request, if supplied
		if len(tagsToCreate) > 0 {
			// Get the existing tags for checking against those on the request
			existingTags, err := as.repository.GetTags(ctx, tagsToCreate)
			if err != nil && err != sql.ErrNoRows {
				return nil, api.NewInternalServerErrorWithContext("tags", err)
			}

			// Roll through the tags on the request to see if we should create any new tags
			for _, tag := range tagsToCreate {
				// If the tag already exists, skip creating it and add it to the list of reference IDs for the article
				if existingTag := findTag(tag, existingTags); existingTag != nil {
					articleTagsToCreate = append(articleTagsToCreate, existingTag.Id)
					continue
				}

				// We've detected a new tag at this point, create it and rollup any errors
				createdTag, err := as.repository.CreateTag(ctx, tag)
				if err != nil {
					return nil, api.NewInternalServerErrorWithContext("tags", err)
				}

				// Add the newly created tag ID to the list to reference from articles
				articleTagsToCreate = append(articleTagsToCreate, createdTag.Id)
			}
		}
	}

	// Create the article, adding the article tags as a text array type in Postgres
	createdArticle, err := as.repository.CreateArticle(ctx, request.UserId, request.Title, articleSlug, request.Description, request.Body, articleTagsToCreate)
	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("articles", err)
	}

	return &domain.ArticleDto{
		Slug:           articleSlug,
		Title:          createdArticle.Title,
		Description:    createdArticle.Description,
		Body:           createdArticle.Body,
		// TagList:        createdArticle.Tags,
		CreatedAt:      createdArticle.CreatedAt,
		UpdatedAt:      createdArticle.UpdatedAt,
		Favorited:      false,
		FavoritesCount: 0,
		Author: domain.AuthorDto{
			Username:  existingUser.Username,
			Bio:       existingUser.Bio,
			Image:     existingUser.Image,
			Following: false,
		},
	}, nil
}

func removeDuplicates(tags *[]string) []string {
	if tags == nil {
		return []string{}
	}

	var depudedTags []string
	{
		for _, tag := range *tags {
			for _, depudedTag := range depudedTags {
				if tag != depudedTag {
					depudedTags = append(depudedTags)
				}
			}
		}
	}

	return depudedTags
}

func containsTag(searchValue string, tags *[]persistence.TagEntity) bool {
	if tags == nil {
		return false
	}

	for _, value := range *tags {
		if value.Tag == searchValue {
			return true
		}
	}

	return false
}

func findTag(searchTag string, tags *[]persistence.TagEntity) *persistence.TagEntity {
	if !containsTag(searchTag, tags) {
		return nil
	}

	for _, tag := range *tags {
		if tag.Tag == searchTag {
			return &tag
		}
	}

	return nil
}
