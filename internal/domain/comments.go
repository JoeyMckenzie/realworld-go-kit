package domain

import (
    "github.com/gofrs/uuid"
    "time"
)

type (
    Comment struct {
        ID        uuid.UUID `json:"id"`
        Body      string    `json:"body"`
        CreatedAt time.Time `json:"createdAt"`
        UpdatedAt time.Time `json:"updatedAt"`
        Author    *Profile  `json:"author"`
    }

    CommentResponse struct {
        Comment *Comment `json:"comment"`
    }

    CommentsResponse struct {
        Comments []Comment `json:"comments"`
    }

    CommentRetrievalRequest struct {
        ID   uuid.UUID
        Slug string `validate:"required"`
    }

    CommentRequest struct {
        Body string `json:"body" validate:"required"`
    }

    CreateCommentRequest struct {
        Comment *CommentRequest `json:"comment" validate:"required"`
        Slug    string          `validate:"-"`
    }
)
