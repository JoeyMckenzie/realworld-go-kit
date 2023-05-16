package comments

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
	"golang.org/x/exp/slog"
)

type (
	CommentsService interface {
		AddComment(ctx context.Context, request domain.CreateCommentRequest, userId uuid.UUID) (*domain.Comment, error)
		DeleteComment(ctx context.Context, userId, commentId uuid.UUID) error
		GetComments(ctx context.Context, request domain.CommentRetrievalRequest, userId uuid.UUID) ([]domain.Comment, error)
	}

	CommentsServiceMiddleware func(next CommentsService) CommentsService

	commentsService struct {
		logger             *slog.Logger
		articlesRepository repositories.ArticlesRepository
		commentsRepository repositories.CommentsRepository
	}
)

func NewCommentsService(logger *slog.Logger, commentsRepository repositories.CommentsRepository, articlesRepository repositories.ArticlesRepository) CommentsService {
	return &commentsService{
		logger:             logger,
		commentsRepository: commentsRepository,
		articlesRepository: articlesRepository,
	}
}
