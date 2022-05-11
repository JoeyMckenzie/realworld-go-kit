package articles

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"
    "time"
)

type (
    UpsertArticleResponse struct {
        Article *ArticleDto `json:"article"`
    }

    ArticleDto struct {
        Slug           string           `json:"slug"`
        Title          string           `json:"title"`
        Description    string           `json:"description"`
        Body           string           `json:"body"`
        TagList        []string         `json:"tagList"`
        CreatedAt      time.Time        `json:"createdAt"`
        UpdatedAt      time.Time        `json:"updatedAt"`
        Favorited      bool             `json:"favorited"`
        FavoritesCount int              `json:"favoritesCount"`
        Author         shared.AuthorDto `json:"author"`
    }

    GetArticlesResponse struct {
        Articles      []*ArticleDto `json:"articles"`
        ArticlesCount int           `json:"articlesCount"`
    }

    GetArticleResponse struct {
        Article *ArticleDto `json:"article"`
    }
)
