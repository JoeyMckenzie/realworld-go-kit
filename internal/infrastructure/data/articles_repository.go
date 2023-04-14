package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	ArticleEntity struct {
		ID          uuid.UUID
		AuthorID    uuid.UUID
		Slug        string
		Title       string
		Description string
		Body        string
		CreatedAt   time.Time `db:"create_at"`
		UpdatedAt   time.Time `db:"update_at"`
	}

	TagEntity struct {
		ID          uuid.UUID
		Description string
		CreatedAt   time.Time `db:"created_at"`
	}

	ArticleTagEntity struct {
		ID        uuid.UUID
		ArticleID uuid.UUID `db:"article_id"`
		TagID     uuid.UUID `db:"tag_id"`
		CreatedAt time.Time `db:"created_at"`
	}

	ArticlesRepository interface {
		BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
		CommitTransaction(tx *sqlx.Tx) error
		CreateArticleWithTags(ctx context.Context, tx *sqlx.Tx, slug, title, description, body, string, tagList []string) (*ArticleEntity, error)
	}

	articlesRepository struct {
		db *sqlx.DB
	}
)

func NewArticlesRepository(db *sqlx.DB) ArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}

func (ar articlesRepository) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return ar.db.BeginTxx(ctx, nil)
}

func (ar articlesRepository) CommitTransaction(tx *sqlx.Tx) error {
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (ar articlesRepository) CreateArticleWithTags(ctx context.Context, tx *sqlx.Tx, slug, title, description, body, string, tagList []string) (*ArticleEntity, error) {
	//TODO implement me
	panic("implement me")
}
