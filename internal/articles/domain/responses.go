package domain

import "time"

type (
	UpsertArticleResponse struct {
		Article *ArticleDto `json:"article"`
	}

	ArticleDto struct {
		Slug           string    `json:"slug"`
		Title          string    `json:"title"`
		Description    string    `json:"description"`
		Body           string    `json:"body"`
		TagList        []string  `json:"tagList"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		Favorited      bool      `json:"favorited"`
		FavoritesCount int       `json:"favoritesCount"`
		Author         AuthorDto `json:"author"`
	}

	AuthorDto struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	}
)
