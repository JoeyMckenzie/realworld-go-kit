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
)

func (request *CreateArticleServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; title: %s; description: %s", request.UserId, request.Title, request.Description)
}
