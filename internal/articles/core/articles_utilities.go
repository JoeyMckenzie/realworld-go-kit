package core

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

func removeDuplicates(tags *[]string) []string {
	if tags == nil || len(*tags) == 0 {
		return []string{}
	}

	var dedupedTags []string
	tagCountMap := make(map[string]int)

	for _, tag := range *tags {
		if _, exists := tagCountMap[tag]; !exists {
			tagCountMap[tag] = 1
		}
	}

	for tag, _ := range tagCountMap {
		dedupedTags = append(dedupedTags, tag)
	}

	return dedupedTags
}

func containsTag(searchValue string, tags *[]persistence.TagEntity) bool {
	if tags == nil || len(*tags) == 0 {
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

func mapArticleEntityToDto(article *persistence.ArticleEntity, user *usersPersistence.UserEntity, tagList []string, favorited, following bool) *domain.ArticleDto {
	return &domain.ArticleDto{
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		TagList:        tagList,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		Favorited:      favorited,
		FavoritesCount: article.Favorites,
		Author: domain.AuthorDto{
			Username:  user.Username,
			Bio:       user.Bio,
			Image:     user.Image,
			Following: following,
		},
	}
}

func handleArticleListing(ctx context.Context, as *articlesServices, articles *[]persistence.ArticleEntity, request *domain.GetArticlesServiceRequest, err error, checkForFavoritedStatus bool) (*[]domain.ArticleDto, error) {
	articlesNotFound := articles == nil || utilities.IsNotFound(err) || len(*articles) == 0

	if utilities.IsValidDbError(err) {
		return nil, api.NewInternalServerErrorWithContext("articles", err)
	} else if articlesNotFound {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "articles", utilities.ErrArticlesNotFound)
	}

	var response []domain.ArticleDto

	for _, article := range *articles {
		mappedArticle, err := makeArticleMapping(ctx, as, article, request.UserId, checkForFavoritedStatus)

		if err != nil {
			return nil, api.NewInternalServerErrorWithContext("articles", err)
		}

		response = append(response, *mappedArticle)
	}

	return &response, nil
}

func makeArticleMapping(ctx context.Context, as *articlesServices, article persistence.ArticleEntity, userId int, checkForFavoritedStatus bool) (*domain.ArticleDto, error) {
	// Get the user that authored the article
	user, err := as.usersRepository.GetUser(ctx, article.UserId)
	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("articles", err)
	}

	// Get any associated tags from the article
	articleTags, err := as.repository.GetArticleTags(ctx, article.Id)
	if utilities.IsValidDbError(err) {
		return nil, api.NewInternalServerErrorWithContext("articles", err)
	}

	// Map the article tags, if any were found
	var mappedArticleTags []string
	if articleTags != nil {
		for _, tag := range *articleTags {
			mappedArticleTags = append(mappedArticleTags, tag)
		}
	}

	// If the user ID is included on the request, check if they've favorited the article or followed the author
	followingUser := false
	if checkForFavoritedStatus {
		if userId > -1 {
			followingUser = as.usersRepository.IsFollowingUser(ctx, userId, article.UserId)
		}
	} else {
		followingUser = true
	}

	favorited := as.repository.UserHasFavoritedArticle(ctx, userId, article.Id)

	return mapArticleEntityToDto(&article, user, mappedArticleTags, favorited, followingUser), nil
}