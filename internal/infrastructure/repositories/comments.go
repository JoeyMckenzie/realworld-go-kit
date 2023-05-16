package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type (
	CommentCompositeQuery struct {
		ID        uuid.UUID `db:"id"`
		Body      string    `db:"body"`
		CreatedAt string    `db:"created_at"`
		UpdatedAt string    `db:"updated_at"`
		Username  string    `db:"author_username"`
		AuthorId  uuid.UUID `db:"author_id"`
		Bio       string    `db:"author_bio"`
		Image     string    `db:"author_image"`
		Following bool      `db:"following_author"`
	}

	CommentsRepository interface {
		AddComment(ctx context.Context, articleId, userId uuid.UUID, comment string) (*CommentCompositeQuery, error)
		DeleteComment(ctx context.Context, commentId uuid.UUID) error
		GetArticleComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (*CommentCompositeQuery, error)
		GetLatestArticleCommentByUser(ctx context.Context, articleId, userId uuid.UUID) (*CommentCompositeQuery, error)
		GetArticleComments(ctx context.Context, userId uuid.UUID, slug string) ([]CommentCompositeQuery, error)
	}

	commentsRepository struct {
		db *sqlx.DB
	}
)

func NewCommentsRepository(db *sqlx.DB) CommentsRepository {
	return &commentsRepository{
		db: db,
	}
}

func (c CommentCompositeQuery) ToComment() (*domain.Comment, error) {
	createdAt, createdTimeParseErr := time.Parse(time.DateTime, c.CreatedAt)
	updatedAt, updatedTimeParseErr := time.Parse(time.DateTime, c.UpdatedAt)

	if err := errors.Join(createdTimeParseErr, updatedTimeParseErr); err != nil {
		return &domain.Comment{}, shared.ErrorWithContext("error while parsing dates for comment", err)
	}

	return &domain.Comment{
		ID:        c.ID,
		Body:      c.Body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Author: &domain.Profile{
			Username:  c.Username,
			Image:     c.Image,
			Bio:       c.Image,
			Following: c.Following,
		},
	}, nil
}

func (r *commentsRepository) AddComment(ctx context.Context, articleId, userId uuid.UUID, comment string) (*CommentCompositeQuery, error) {
	const sql = `
INSERT INTO comments (id, author_id, article_id, body)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), UUID_TO_BIN(?), ?)`

	if _, err := r.db.ExecContext(ctx, sql, userId, articleId, comment); err != nil {
		return nil, err
	}

	return r.GetLatestArticleCommentByUser(ctx, articleId, userId)
}

func (r *commentsRepository) DeleteComment(ctx context.Context, commentId uuid.UUID) error {
	const sql = "DELETE FROM articles WHERE id = UUID_TO_BIN(?)"

	if _, err := r.db.ExecContext(ctx, sql, commentId); err != nil {
		return err
	}

	return nil
}

func (r *commentsRepository) GetArticleComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (*CommentCompositeQuery, error) {
	var comment CommentCompositeQuery
	{
		const sql = `
SELECT c.id                                                                     AS "id",
       c.body                                                                   AS "body",
       c.created_at                                                             AS "created_at",
       c.updated_at                                                             AS "updated_at",
       c.author_id                                                              AS "author_id",
       u.id                                                                     AS "author_username",
       u.username                                                               AS "author_username",
       u.bio                                                                    AS "author_bio",
       u.image                                                                  AS "author_image",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE (followee_id, follower_id) = (c.author_id, UUID_TO_BIN(?))) AS "following_author"
FROM comments c
     LEFT JOIN users u ON c.author_id = u.id
WHERE c.id = UUID_TO_BIN(?)`

		if err := r.db.GetContext(ctx, &comment, sql, userId, commentId); err != nil {
			return nil, err
		}
	}

	return &comment, nil
}

func (r *commentsRepository) GetLatestArticleCommentByUser(ctx context.Context, articleId, userId uuid.UUID) (*CommentCompositeQuery, error) {
	var comment CommentCompositeQuery
	{
		const sql = `
SELECT c.id                                                                     AS "id",
       c.body                                                                   AS "body",
       c.created_at                                                             AS "created_at",
       c.updated_at                                                             AS "updated_at",
       u.id                                                                     AS "author_username",
       u.username                                                               AS "author_username",
       u.bio                                                                    AS "author_bio",
       u.image                                                                  AS "author_image",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE (followee_id, follower_id) = (c.author_id, UUID_TO_BIN(?))) AS "following_author"
FROM comments c
     JOIN users u ON c.author_id = u.id
WHERE (c.author_id, c.article_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))
ORDER BY created_at DESC
LIMIT 1`

		if err := r.db.GetContext(ctx, &comment, sql, userId, userId, articleId); err != nil {
			return nil, err
		}
	}

	return &comment, nil
}

func (r *commentsRepository) GetArticleComments(ctx context.Context, userId uuid.UUID, slug string) ([]CommentCompositeQuery, error) {
	var comments []CommentCompositeQuery
	{
		const sql = `
SELECT c.id                                                                     AS "id",
       c.body                                                                   AS "body",
       c.created_at                                                             AS "created_at",
       c.updated_at                                                             AS "updated_at",
       u.id                                                                     AS "author_username",
       u.username                                                               AS "author_username",
       u.bio                                                                    AS "author_bio",
       u.image                                                                  AS "author_image",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE (followee_id, follower_id) = (c.author_id, UUID_TO_BIN(?))) AS "following_author"
FROM comments c
    JOIN users u ON c.author_id = u.id
    JOIN articles a ON c.article_id = a.id
WHERE a.slug = ?
ORDER BY created_at DESC`

		if err := r.db.SelectContext(ctx, &comments, sql, userId, slug); err != nil {
			return nil, err
		}
	}

	return comments, nil
}
