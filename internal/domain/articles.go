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
        Author         *Profile
    }

    ArticleResponse struct {
        Article `json:"article"`
    }

    CreateArticleRequest struct {
        Title       string   `json:"title"`
        Description string   `json:"description"`
        Body        string   `json:"body"`
        TagList     []string `json:"tagList"`
    }
)
