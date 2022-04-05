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
		Tags        []string  `db:"tags"`
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
		Tag       string    `db:"tag"`
	}

	ArticlesRepository interface {
		FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error)
		GetArticles(ctx context.Context, tag, author, favorited string, limit, offset int) (*[]ArticleEntity, error)
		CreateArticle(ctx context.Context, userId int, title, slug, description, body string, tagList []int) (*ArticleEntity, error)
		CreateTag(ctx context.Context, tag string) (*TagEntity, error)
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

func (ar *articlesRepository) GetArticles(ctx context.Context, tag, author, favorited string, limit, offset int) (*[]ArticleEntity, error) {
	var articles []ArticleEntity
	var sql strings.Builder

	const sqlTest = `
SELECT a.id,
       a.created_at,
       a.updated_at,
       a.title,
       a.slug,
       a.description,
       a.body
FROM public.articles a
JOIN article_tags at on a.id = at.article_id
JOIN tags t on at.tag_id = t.id
WHERE t.tag = $1
`

	sql.WriteString(`
SELECT *
FROM public.articles a
LEFT JOIN public.article_tags at on a.id = at.article_id`)

	if tag != "" {
		sql.WriteString(`
JOIN tags t on at.tag_id = t.id
WHERE t.tag = $1`)
	}

	if err := ar.db.SelectContext(ctx, articles, sqlTest, tag, author, favorited, limit, offset); err != nil {
		return nil, err
	}

	return nil, nil
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, userId int, title, slug, description, body string, tagList []int) (*ArticleEntity, error) {
	var article ArticleEntity

	const sql = `
INSERT INTO public.articles (created_at, updated_at, title, slug, description, body, user_id, tags)
VALUES (current_timestamp, current_timestamp, $1::VARCHAR, $2::VARCHAR, $3::VARCHAR, $4::VARCHAR, $5::INTEGER, $6::INTEGER[])
RETURNING *`

	if err := ar.db.GetContext(ctx, &article, sql, title, slug, description, body, userId, tagList); err != nil {
		return nil, err
	}

	return &article, nil
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

func (ar *articlesRepository) CreateTag(ctx context.Context, tag string) (*TagEntity, error) {
	var createdTag TagEntity

	const sql = `
INSERT INTO tags (created_at, tag)
VALUES (current_timestamp, $1::VARCHAR)
RETURNING *`

	if err := ar.db.GetContext(ctx, &tag, sql, tag); err != nil {
		return nil, err
	}

	return &createdTag, nil
}

func (ar *articlesRepository) GetTags(ctx context.Context, tags []string) (*[]TagEntity, error) {
	var articleTags []TagEntity
	var sql string
	{
		if len(tags) > 0 {
			sql = `
SELECT *
FROM public.tags
WHERE tag IN ($1::VARCHAR)`
		} else {
			sql = `
SELECT *
FROM public.tags`
		}
	}

	if err := ar.db.SelectContext(ctx, &articleTags, sql, strings.Join(tags, ",")); err != nil {
		return nil, err
	}

	return &articleTags, nil
}
