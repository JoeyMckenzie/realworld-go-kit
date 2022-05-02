package domain

import (
	sharedDomain "github.com/joeymckenzie/realworld-go-kit/internal/shared/domain"
	"time"
)

type (
	CommentDto struct {
		Id        int                    `json:"id"`
		CreatedAt time.Time              `json:"createdAt"`
		UpdatedAt time.Time              `json:"updatedAt"`
		Body      string                 `json:"body"`
		Author    sharedDomain.AuthorDto `json:"author"`
	}

	CommentResponse struct {
		Comment *CommentDto `json:"comment"`
	}

	CommentsResponse struct {
		Comments []*CommentDto `json:"comments"`
	}
)
