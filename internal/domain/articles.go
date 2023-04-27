package domain

import (
    "time"
)

type (
    Article struct {
        Slug           string    `json:"slug"`
        Title          string    `json:"title"`
        Description    string    `json:"description"`
        Body           string    `json:"body"`
        TagList        []string  `json:"tagList"`
        CreatedAt      time.Time `json:"createdAt"`
        UpdatedAt      time.Time `json:"updatedAt"`
        Favorited      bool      `json:"favorited"`
        FavoritesCount int       `json:"favoritesCount"`
        Author         *Profile  `json:"author"`
    }

    ArticleResponse struct {
        Article *Article `json:"article"`
    }

    ArticlesResponse struct {
        Articles      []Article `json:"articles"`
        ArticlesCount int       `json:"articlesCount"`
    }

    ArticleRequest struct {
        Title       string   `json:"title" validate:"required"`
        Description string   `json:"description" validate:"required"`
        Body        string   `json:"body" validate:"required"`
        TagList     []string `json:"tagList"`
    }

    CreateArticleRequest struct {
        Article *ArticleRequest `json:"article" validate:"required"`
    }

    ListArticlesRequest struct {
        Limit     int
        Offset    int
        Tag       string
        Author    string
        Favorited string
    }
)
