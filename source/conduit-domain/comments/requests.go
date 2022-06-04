package comments

import "fmt"

type (
	AddCommentApiRequest struct {
		Slug    string
		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`
	}

	AddArticleCommentServiceRequest struct {
		Body   string `validate:"required"`
		UserId int    `validate:"required"`
		Slug   string `validate:"required"`
	}

	DeleteArticleCommentServiceRequest struct {
		CommentId int `validate:"required"`
		UserId    int `validate:"required"`
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

	return fmt.Sprintf("userId: %d; commentId: %d", r.UserId, r.CommentId)
}

func (r *AddArticleCommentServiceRequest) ToSafeLoggingStruct() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; slug: %s; body: %s", r.UserId, r.Slug, r.Body)
}
