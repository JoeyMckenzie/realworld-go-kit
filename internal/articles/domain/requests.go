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
		Slug   string
	}
)

func (request *UpdateArticleServiceRequest) ToSafeLoggingStruct() string {
	//TODO implement me
	panic("implement me")
}

func (request *GetArticleServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s", request.UserId, request.Slug)
}

func (request *GetArticlesServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("tag: %s; author: %s; favoried: %s; limit: %d; offset: %d", request.Tag, request.Author, request.Favorited, request.Limit, request.Offset)
}

func (request *CreateArticleServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; title: %s; description: %s", request.UserId, request.Title, request.Description)
}
