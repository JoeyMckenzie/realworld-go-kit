package domain

import "fmt"

type (
	UpsertArticleApiRequest struct {
		Article UpsertArticleDto `json:"article"`
	}

	UpsertArticleDto struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Body        string    `json:"body"`
		TagList     *[]string `json:"tagList,omitempty"`
	}

	CreateArticleServiceRequest struct {
		UserId      int    `validate:"required"`
		Title       string `validate:"required"`
		Description string `validate:"required"`
		Body        string `validate:"required"`
		TagList     *[]string
	}

	GetArticlesServiceRequest struct {
		Tag       string
		Author    string
		Favorited string
		Limit     int
		Offset    int
	}
)

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
