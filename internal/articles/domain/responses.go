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

	GetArticlesResponse struct {
		Articles      []*ArticleDto `json:"articles"`
		ArticlesCount int           `json:"articlesCount"`
	}

	GetArticleResponse struct {
		Article *ArticleDto `json:"article"`
	}

	GetTagsResponse struct {
		Tags []string `json:"tags"`
	}

	CommentDto struct {
		Id        int       `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Body      string    `json:"body"`
		Author    AuthorDto `json:"author"`
	}

	CommentResponse struct {
		Comment *CommentDto `json:"comment"`
	}

	CommentsResponse struct {
		Comments []*CommentDto `json:"comments"`
	}
)
