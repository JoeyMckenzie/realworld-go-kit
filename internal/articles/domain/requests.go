package domain

import "fmt"

type (
	CreateArticleApiRequest struct {
		Article CreateArticleDto `json:"article"`
	}

	UpdateArticleApiRequest struct {
		Article UpdateArticleDto `json:"article"`
	}

	CreateArticleDto struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Body        string    `json:"body"`
		TagList     *[]string `json:"tagList,omitempty"`
	}

	UpdateArticleDto struct {
		Slug        string
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		Body        *string `json:"body,omitempty"`
	}

	CreateArticleServiceRequest struct {
		ArticleSlug string
		UserId      int    `validate:"required"`
		Title       string `validate:"required"`
		Description string `validate:"required"`
		Body        string `validate:"required"`
		TagList     *[]string
	}

	UpdateArticleServiceRequest struct {
		ArticleSlug string `validate:"required"`
		UserId      int    `validate:"required"`
		Title       *string
		Description *string
		Body        *string
	}

	DeleteArticleServiceRequest struct {
		ArticleSlug string `validate:"required"`
		UserId      int    `validate:"required"`
	}

	AddArticleCommentServiceRequest struct {
		Body   string `validate:"required"`
		UserId int    `validate:"required"`
		Slug   string `validate:"required"`
	}

	DeleteArticleCommentServiceRequest struct {
		CommentId int    `validate:"required"`
		UserId    int    `validate:"required"`
		Slug      string `validate:"required"`
	}

	GetArticlesServiceRequest struct {
		UserId    int
		Tag       string
		Author    string
		Favorited string
		Limit     int
		Offset    int
	}

	GetArticleServiceRequest struct {
		UserId int
		Slug   string `validate:"required"`
	}

	ArticleFavoriteServiceRequest struct {
		UserId int    `validate:"required"`
		Slug   string `validate:"required"`
	}

	GetCommentsServiceRequest struct {
		UserId int
		Slug   string `validate:"required"`
	}
)

func (r *GetCommentsServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s", r.UserId, r.Slug)
}

func (r *DeleteArticleCommentServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s; commentId: %d", r.UserId, r.Slug, r.CommentId)
}

func (r *AddArticleCommentServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s; body: %s", r.UserId, r.Slug, r.Body)
}

func (r *ArticleFavoriteServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s", r.UserId, r.Slug)
}

func (r *DeleteArticleServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s", r.UserId, r.ArticleSlug)
}

func (r *UpdateArticleServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s; title: %v; description: %v; body: %v", r.UserId, r.ArticleSlug, r.Title, r.Description, r.Body)
}

func (r *GetArticleServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s", r.UserId, r.Slug)
}

func (r *GetArticlesServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("tag: %s; author: %s; favorited: %s; limit: %d; offset: %d", r.Tag, r.Author, r.Favorited, r.Limit, r.Offset)
}

func (r *CreateArticleServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; title: %s; description: %s", r.UserId, r.Title, r.Description)
}
