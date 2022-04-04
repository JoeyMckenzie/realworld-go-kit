package persistence

import (
	"context"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type (
	ArticleEntity struct {
		Id          int       `db:"id"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		Title       string    `db:"title"`
		Slug        string    `db:"slug"`
		Description string    `db:"description"`
		Body        string    `db:"body"`
		UserId      int       `db:"user_id"`
	}

	ArticleTagEntity struct {
		Id        int       `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		TagId     int       `db:"tag"`
		ArticleId int       `db:"article_id"`
	}

	TagEntity struct {
		Id        int       `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		Tag       string    `db:"tag"`
	}

	ArticlesRepository interface {
		FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error)
		CreateArticle(ctx context.Context, userId int, title, slug, description, body string) (*ArticleEntity, error)
		CreateArticleTag(ctx context.Context, tagId int, articleId int) (*ArticleTagEntity, error)
		GetArticleTags(ctx context.Context, tags []string) (*[]ArticleTagEntity, error)
		CreateTag(ctx context.Context, tag string) (*TagEntity, error)
		GetTag(ctx context.Context, tag string) (*TagEntity, error)
		GetTags(ctx context.Context, tags []string) (*[]TagEntity, error)
	}

	articlesRepository struct {
		db *sqlx.DB
	}

	ArticlesRepositoryMiddleware func(next ArticlesRepository) ArticlesRepository
)

func NewArticlesRepository(db *sqlx.DB) ArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, userId int, title, slug, description, body string) (*ArticleEntity, error) {
	var article ArticleEntity

	const sql = `
INSERT INTO public.articles (created_at, updated_at, title, slug, description, body, user_id)
VALUES (current_timestamp, current_timestamp, $1::VARCHAR, $2::VARCHAR, $3::VARCHAR, $4::VARCHAR, $5::INTEGER)
RETURNING *`

	if err := ar.db.GetContext(ctx, &article, sql, title, slug, description, body, userId); err != nil {
		return nil, err
	}

	return &article, nil
}

func (ar *articlesRepository) CreateArticleTag(ctx context.Context, tagId int, articleId int) (*ArticleTagEntity, error) {
	var articleTag ArticleTagEntity

	const sql = `
INSERT INTO public.article_tags (created_at, updated_at, tag_id, article_id)
VALUES (current_timestamp, current_timestamp, $1::INTEGER, $2::INTEGER)
RETURNING *`

	if err := ar.db.GetContext(ctx, &articleTag, sql, tagId, articleId); err != nil {
		return nil, err
	}

	return &articleTag, nil
}

func (ar *articlesRepository) FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error) {
	var article ArticleEntity

	const sql = `
SELECT *
FROM public.articles
WHERE slug = $1::VARCHAR`

	if err := ar.db.GetContext(ctx, &article, sql, slug); err != nil {
		return nil, err
	}

	return &article, nil
}

func (ar *articlesRepository) GetTags(ctx context.Context, tags []string) (*[]TagEntity, error) {
	var articleTags []TagEntity

	const sql = `
SELECT *
FROM public.tags
WHERE tag IN ($1::VARCHAR)`

	if err := ar.db.SelectContext(ctx, &articleTags, sql, strings.Join(tags, ",")); err != nil {
		return nil, err
	}

	return &articleTags, nil
}

func (ar *articlesRepository) CreateTag(ctx context.Context, tag string) (*TagEntity, error) {
	var articleTag TagEntity

	const sql = `
INSERT INTO public.tags (created_at, updated_at, tag)
VALUES (current_timestamp, current_timestamp, $1::VARCHAR)
RETURNING *`

	if err := ar.db.GetContext(ctx, &articleTag, sql, tag); err != nil {
		return nil, err
	}

	return &articleTag, nil
}

func (ar *articlesRepository) GetArticleTags(ctx context.Context, tags []string) (*[]ArticleTagEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (ar *articlesRepository) GetTag(ctx context.Context, tag string) (*TagEntity, error) {
	var articleTag TagEntity

	const sql = `
SELECT *
FROM public.tags
WHERE tag = $1::VARCHAR`
	if err := ar.db.GetContext(ctx, &articleTag, sql, tag); err != nil {
		return nil, err
	}

	return &articleTag, nil
}
