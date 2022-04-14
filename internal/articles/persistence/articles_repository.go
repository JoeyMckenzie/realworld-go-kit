package persistence

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
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
		TagId     int       `db:"tag_id"`
		ArticleId int       `db:"article_id"`
	}

	TagEntity struct {
		Id        int       `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		Tag       string    `db:"tag"`
	}

	ArticlesRepository interface {
		FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error)
		GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (*[]ArticleEntity, error)
		CreateArticle(ctx context.Context, userId int, title, slug, description, body string) (*ArticleEntity, error)
		CreateTag(ctx context.Context, tag string) (*TagEntity, error)
		GetTags(ctx context.Context, tags []string) (*[]TagEntity, error)
		CreateArticleTag(ctx context.Context, tagId, articleId int) (*ArticleTagEntity, error)
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

func (ar *articlesRepository) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (*[]ArticleEntity, error) {
	var articles []ArticleEntity
	var parameterizedInput []interface{}
	var sql strings.Builder
	{
		sql.WriteString(`
select a.id,
       a.created_at,
       a.updated_at,
       a.title,
       a.slug,
       a.description,
       a.body
	   a.user_id
from public.articles a`)
	}

	if request.Tag != "" {
		parameterizedInput = append(parameterizedInput, request.Tag)
		sql.WriteString(fmt.Sprintf(`
join public.article_tags at on at.article_id = a.id
join public.tags t on t.id = at.tag_id
where t.tag = %d::varchar`, parameterizedInput))
	}

	if request.Author != "" {
		parameterizedInput = append(parameterizedInput, request.Author)
		sql.WriteString(fmt.Sprintf(`
join public.users u on u.id = a.user_id
where u.username = %d::varchar`, parameterizedInput))
	}

	if request.Favorited != "" {
		parameterizedInput = append(parameterizedInput, request.Favorited)
		sql.WriteString(fmt.Sprintf(`
join public.users uu on uu.id = a.user_id
where uu.username = %d::varchar`, parameterizedInput))
	}

	parameterizedInput = append(parameterizedInput, request.Limit)
	parameterizedInput = append(parameterizedInput, request.Offset)
	sql.WriteString(fmt.Sprintf(`
limit $%d::integer
offset $%d::integer
`, len(parameterizedInput)-1, len(parameterizedInput)))

	getArticlesSql := sql.String()

	if err := ar.db.SelectContext(ctx, &articles, getArticlesSql, parameterizedInput...); err != nil {
		return nil, err
	}

	return &articles, nil
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, userId int, title, slug, description, body string) (*ArticleEntity, error) {
	var article ArticleEntity

	const sql = `
insert into public.articles (created_at, updated_at, title, slug, description, body, user_id)
values (current_timestamp, current_timestamp, $1::varchar, $2::varchar, $3::varchar, $4::varchar, $5::bigint)
returning *`

	if err := ar.db.GetContext(ctx, &article, sql, title, slug, description, body, userId); err != nil {
		return nil, err
	}

	return &article, nil
}

func (ar *articlesRepository) FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error) {
	var article ArticleEntity

	const sql = `
select *
from public.articles
where slug = $1::varchar`

	if err := ar.db.GetContext(ctx, &article, sql, slug); err != nil {
		return nil, err
	}

	return &article, nil
}

func (ar *articlesRepository) CreateTag(ctx context.Context, tag string) (*TagEntity, error) {
	var createdTag TagEntity

	const sql = `
insert into tags (created_at, tag)
values (current_timestamp, $1::varchar)
returning *`

	if err := ar.db.GetContext(ctx, &createdTag, sql, tag); err != nil {
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
select *
from public.tags
where tag in ($1::varchar)`
		} else {
			sql = `
select *
from public.tags`
		}
	}

	if err := ar.db.SelectContext(ctx, &articleTags, sql, strings.Join(tags, ",")); err != nil {
		return nil, err
	}

	return &articleTags, nil
}

func (ar *articlesRepository) CreateArticleTag(ctx context.Context, tagId, articleId int) (*ArticleTagEntity, error) {
	var createdArticleTag ArticleTagEntity

	const sql = `
insert into article_tags (created_at, tag_id, article_id)
values (current_timestamp, $1::bigint, $2::bigint)
returning *`

	if err := ar.db.GetContext(ctx, &createdArticleTag, sql, tagId, articleId); err != nil {
		return nil, err
	}

	return &createdArticleTag, nil
}
