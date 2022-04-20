package core

import (
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
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

func containsTag(searchTag string, tags []*ent.Tag) bool {
	if len(tags) == 0 {
		return false
	}

	for _, tag := range tags {
		if tag != nil && tag.Tag == searchTag {
			return true
		}
	}

	return false
}

func firstOrDefaultTag(searchTag string, tags []*ent.Tag) *ent.Tag {
	if !containsTag(searchTag, tags) {
		return nil
	}

	for _, tag := range tags {
		if tag.Tag == searchTag {
			return tag
		}
	}

	return nil
}

func mapArticleEntityToDto(article *ent.Article, user *ent.User, tagList []string, favorited, following bool) *domain.ArticleDto {
	return &domain.ArticleDto{
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		TagList:     tagList,
		CreatedAt:   article.CreateTime,
		UpdatedAt:   article.UpdateTime,
		Favorited:   favorited,
		Author: domain.AuthorDto{
			Username:  user.Username,
			Bio:       user.Bio,
			Image:     user.Image,
			Following: following,
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
		userHasFAvo := false

		if request.UserId > 0 {
			if !userHasFavorited {
				for _, articleFavorite := range queriedArticle.Edges.Favorites {
					if articleFavorite.UserID == request.UserId {
						userHasFavorited = true
						break
					}
				}
			}

			for _, userFollower := range queriedArticle.Edges.Author.Edges.Followees {
				if userFollower.FollowerID == request.UserId {
					userHasFAvo = true
					break
				}
			}
		}

		mappedArticle := &domain.ArticleDto{
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
				Following: userHasFAvo,
			},
		}

		mappedArticles = append(mappedArticles, mappedArticle)
	}

	return mappedArticles, nil
}
