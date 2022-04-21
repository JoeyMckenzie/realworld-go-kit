package core

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
)

func removeDuplicates(tags *[]string) []string {
	if tags == nil || len(*tags) == 0 {
		return []string{}
	}

	var dedupedTags []string
	tagCountMap := make(map[string]int)

	for _, tagsToCreate := range *tags {
		if _, exists := tagCountMap[tagsToCreate]; !exists {
			tagCountMap[tagsToCreate] = 1
		}
	}

	for tagKey, _ := range tagCountMap {
		dedupedTags = append(dedupedTags, tagKey)
	}

	return dedupedTags
}

func containsTag(searchTag string, tags []*ent.Tag) bool {
	if len(tags) == 0 {
		return false
	}

	for _, tagValue := range tags {
		if tagValue != nil && tagValue.Tag == searchTag {
			return true
		}
	}

	return false
}

func firstOrDefaultTag(searchTag string, tags []*ent.Tag) *ent.Tag {
	if !containsTag(searchTag, tags) {
		return nil
	}

	for _, availableTags := range tags {
		if availableTags.Tag == searchTag {
			return availableTags
		}
	}

	return nil
}

func makeArticleMapping(queriedArticle *ent.Article, defaultHasFavorited bool, userId int) *domain.ArticleDto {
	var tags []string
	{
		// Aggregate all the tags on the article
		for _, articleTag := range queriedArticle.Edges.ArticleTags {
			if articleTag.Edges.Tag != nil {
				tags = append(tags, articleTag.Edges.Tag.Tag)
			}
		}
	}

	// Set default return values for following and favorited
	// In the case of favoriting, if we're calling from the feed context, we can bypass
	// iteration of the article favorites, as we already know the user has favorited the article
	userHasFavorited := defaultHasFavorited
	userIsFollowing := false

	if userId > 0 {
		if !userHasFavorited {
			for _, articleFavorite := range queriedArticle.Edges.Favorites {
				if articleFavorite.UserID == userId {
					userHasFavorited = true
					break
				}
			}
		}

		for _, userFollower := range queriedArticle.Edges.Author.Edges.Followees {
			if userFollower.FollowerID == userId {
				userIsFollowing = true
				break
			}
		}
	}

	return &domain.ArticleDto{
		Slug:           queriedArticle.Slug,
		Title:          queriedArticle.Title,
		Description:    queriedArticle.Description,
		Body:           queriedArticle.Body,
		TagList:        tags,
		CreatedAt:      queriedArticle.CreateTime,
		UpdatedAt:      queriedArticle.UpdateTime,
		Favorited:      userHasFavorited,
		FavoritesCount: len(queriedArticle.Edges.Favorites),
		Author: domain.AuthorDto{
			Username:  queriedArticle.Edges.Author.Username,
			Bio:       queriedArticle.Edges.Author.Bio,
			Image:     queriedArticle.Edges.Author.Image,
			Following: userIsFollowing,
		},
	}
}

func makeArticlesMapping(articles []*ent.Article, request *domain.GetArticlesServiceRequest, defaultHasFavorited bool, err error) ([]*domain.ArticleDto, error) {
	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("articles", err)
	} else if len(articles) == 0 {
		return defaultArticlesResponse, nil
	}

	var mappedArticles []*domain.ArticleDto

	for _, queriedArticle := range articles {
		mappedArticle := makeArticleMapping(queriedArticle, defaultHasFavorited, request.UserId)
		mappedArticles = append(mappedArticles, mappedArticle)
	}

	return mappedArticles, nil
}

func (as *articlesService) makeArticleTagsMapping(ctx context.Context, transaction *ent.Tx, tagsToCreate []string) ([]*ent.ArticleTagCreate, error) {
	var articleTagsToBulkInsert []*ent.ArticleTagCreate

	if len(tagsToCreate) > 0 {

		// Get the existing tags for checking against those on the request
		existingTags, err := as.client.Tag.
			Query().
			Where(tag.TagIn(tagsToCreate...)).
			All(ctx)

		if err != nil {
			return nil, api.NewInternalServerErrorWithContext("tags", err)
		}

		// Roll through the tags on the request to see if we should create any new tags
		for _, tagToCreate := range tagsToCreate {
			articleTag := as.client.ArticleTag.Create()

			// If the tagToCreate already exists, skip creating it and add it to the list of reference IDs for the article
			if existingTag := firstOrDefaultTag(tagToCreate, existingTags); existingTag != nil {
				articleTag.SetTagID(existingTag.ID)
			} else {
				// We've detected a new tag to create at this point, append to the bulk insert list
				createdTag, err := transaction.Tag.
					Create().
					SetTag(tagToCreate).
					Save(ctx)

				if err != nil {
					_ = transaction.Rollback()
					return nil, api.NewInternalServerErrorWithContext("tags", err)
				}

				articleTag.SetTagID(createdTag.ID)
			}

			articleTagsToBulkInsert = append(articleTagsToBulkInsert, articleTag)
		}
	}

	return articleTagsToBulkInsert, nil
}
