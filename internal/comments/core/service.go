package core

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/comment"
	"github.com/joeymckenzie/realworld-go-kit/internal/comments/domain"
	sharedDomain "github.com/joeymckenzie/realworld-go-kit/internal/shared/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
)

type (
	CommentsService interface {
		AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (*domain.CommentDto, error)
		DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) error
		GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) ([]*domain.CommentDto, error)
	}

	commentsService struct {
		validator *validator.Validate
		client    *ent.Client
	}

	CommentsServiceMiddleware func(service CommentsService) CommentsService
)

func NewCommentsService(validator *validator.Validate, client *ent.Client) CommentsService {
	return &commentsService{
		validator: validator,
		client:    client,
	}
}

func (cs *commentsService) AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (*domain.CommentDto, error) {
	// Verify the user and article exists before adding the comment
	existingArticle, err := cs.client.Article.
		Query().
		Where(article.Slug(request.Slug)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "article", utilities.ErrArticlesNotFound)
	}

	existingUser, err := cs.client.User.Get(ctx, request.UserId)

	if ent.IsNotFound(err) {
		return nil, api.NewApiErrorWithContext(http.StatusNotFound, "article", utilities.ErrUserNotFound)
	}

	// Create the comment, then add it to the article
	newComment, err := cs.client.Comment.
		Create().
		SetBody(request.Body).
		SetArticle(existingArticle).
		SetUser(existingUser).
		Save(ctx)

	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("comment", err)
	}

	return &domain.CommentDto{
		Id:        newComment.ID,
		CreatedAt: newComment.CreateTime,
		UpdatedAt: newComment.UpdateTime,
		Body:      newComment.Body,
		Author: sharedDomain.AuthorDto{
			Username:  existingUser.Username,
			Bio:       existingUser.Bio,
			Image:     existingUser.Image,
			Following: false,
		},
	}, nil
}

func (cs *commentsService) DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) error {
	// Verify the comment with matching IDs exists before attempting to remove
	// In theory, retrieving the comment by ID before deleting should suffice,
	// but we should be wary of malicious users trying to delete comments that
	// don't belong to them by spoofing their user ID on the request
	existingComment, err := cs.client.Comment.
		Query().
		Where(
			comment.ID(request.CommentId),
			comment.UserID(request.UserId),
		).
		First(ctx)

	if ent.IsNotFound(err) {
		return api.NewApiErrorWithContext(http.StatusNotFound, "article", utilities.ErrCommentNotFound)
	}

	if err = cs.client.Comment.DeleteOne(existingComment).Exec(ctx); err != nil {
		return api.NewInternalServerErrorWithContext("comment", err)
	}

	return nil
}

func (cs *commentsService) GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) ([]*domain.CommentDto, error) {
	existingComments, err := cs.client.Comment.
		Query().
		WithUser(func(query *ent.UserQuery) {
			query.WithFollowees()
		}).
		All(ctx)

	if err != nil {
		return nil, api.NewInternalServerErrorWithContext("comments", err)
	}

	var mappedComments []*domain.CommentDto

	// Iterate over the articles to determine if the requesting user is following the authors of said comment
	for _, existingComment := range existingComments {
		mappedComment := &domain.CommentDto{
			Id:        existingComment.ID,
			CreatedAt: existingComment.CreateTime,
			UpdatedAt: existingComment.UpdateTime,
			Body:      existingComment.Body,
			Author: sharedDomain.AuthorDto{
				Username:  existingComment.Edges.User.Username,
				Bio:       existingComment.Edges.User.Bio,
				Image:     existingComment.Edges.User.Image,
				Following: false,
			},
		}

		for _, follower := range existingComment.Edges.User.Edges.Followers {
			if follower.FolloweeID == request.UserId {
				mappedComment.Author.Following = true
			}
		}

		mappedComments = append(mappedComments, mappedComment)
	}

	return mappedComments, nil
}
