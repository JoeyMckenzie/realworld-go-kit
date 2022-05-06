package domain

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-api/shared"
    "time"
)

type (
    CommentDto struct {
        Id        int              `json:"id"`
        CreatedAt time.Time        `json:"createdAt"`
        UpdatedAt time.Time        `json:"updatedAt"`
        Body      string           `json:"body"`
        Author    shared.AuthorDto `json:"author"`
    }

    CommentResponse struct {
        Comment *CommentDto `json:"comment"`
    }

    CommentsResponse struct {
        Comments []*CommentDto `json:"comments"`
    }
)
