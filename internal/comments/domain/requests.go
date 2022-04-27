package domain

import "fmt"

type (
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
