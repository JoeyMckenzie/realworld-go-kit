package core

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
)

type (
	CommentsService interface {
		AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (*domain.CommentDto, error)
		DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) error
		GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) ([]*domain.CommentDto, error)
	}

	commentsService struct {
		validator       *validator.Validate
		client          *ent.Client
		articlesService ArticlesService
	}

	CommentsServiceMiddleware func(service CommentsService) CommentsService
)

func NewCommentsService(validator *validator.Validate, client *ent.Client, articlesService ArticlesService) CommentsService {
	return &commentsService{
		validator:       validator,
		client:          client,
		articlesService: articlesService,
	}
}

func (cs *commentsService) AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (*domain.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *commentsService) DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) error {
	//TODO implement me
	panic("implement me")
}

func (cs *commentsService) GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) ([]*domain.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}
