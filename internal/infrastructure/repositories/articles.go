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
	ArticleCompositeQuery struct {
		ID              uuid.UUID `db:"id"`
		CreatedAt       string    `db:"created_at"`
		UpdatedAt       string    `db:"updated_at"`
		Title           string    `db:"title"`
		Body            string    `db:"body"`
		Description     string    `db:"description"`
		Slug            string    `db:"slug"`
		UserID          uuid.UUID `db:"user_id"`
		Favorited       bool      `db:"favorited"`
		Favorites       int       `db:"favorites"`
		FollowingAuthor bool      `db:"following_author"`
		Username        string    `db:"author_username"`
		Bio             string    `db:"author_bio"`
		Image           string    `db:"author_image"`
		Email           string    `db:"author_email"`
	}

	ArticlesRepository interface {
		BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
		CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleCompositeQuery, error)
		UpdateArticle(ctx context.Context, articleId, userId uuid.UUID, slug, title, description, body string) (*ArticleCompositeQuery, error)
		DeleteArticle(ctx context.Context, articleId uuid.UUID) error
		GetArticles(ctx context.Context, userId uuid.UUID, tag, author, favorited string, limit, offset int) ([]ArticleCompositeQuery, error)
		GetArticle(ctx context.Context, tx *sqlx.Tx, slug string, userId uuid.UUID) (*ArticleCompositeQuery, error)
		AddFavorite(ctx context.Context, articleId, userId uuid.UUID) error
		DeleteFavorite(ctx context.Context, articleId, userId uuid.UUID) error
	}

	articlesRepository struct {
		db *sqlx.DB
	}
)

func (a ArticleCompositeQuery) ToArticle(tagList []string) (*domain.Article, error) {
	createdAt, createdTimeParseErr := time.Parse(time.DateTime, a.CreatedAt)
	updatedAt, updatedTimeParseErr := time.Parse(time.DateTime, a.UpdatedAt)

	if err := errors.Join(createdTimeParseErr, updatedTimeParseErr); err != nil {
		return &domain.Article{}, shared.ErrorWithContext("error while parsing dates for article", err)
	}

	return &domain.Article{
		Slug:           a.Slug,
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		TagList:        tagList,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Favorited:      a.Favorited,
		FavoritesCount: a.Favorites,
		Author: &domain.Profile{
			Username:  a.Username,
			Image:     a.Image,
			Bio:       a.Bio,
			Following: a.FollowingAuthor,
		},
	}, nil
}

func NewArticlesRepository(db *sqlx.DB) ArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}

func (ar *articlesRepository) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return ar.db.BeginTxx(ctx, nil)
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleCompositeQuery, error) {
	const sql = `
INSERT INTO articles (id, author_id, slug, title, description, body)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), ?, ?, ?, ?)`

	if _, err := tx.ExecContext(ctx, sql, authorId, slug, title, description, body); err != nil {
		return &ArticleCompositeQuery{}, err
	}

	return ar.GetArticle(ctx, tx, slug, authorId)
}

func (ar *articlesRepository) UpdateArticle(ctx context.Context, articleId, userId uuid.UUID, slug, title, description, body string) (*ArticleCompositeQuery, error) {
	const sql = `
UPDATE articles
SET title = ?,
    description = ?,
    body = ?
WHERE id = UUID_TO_BIN(?)`

	if _, err := ar.db.ExecContext(ctx, sql, title, description, body, articleId); err != nil {
		return &ArticleCompositeQuery{}, err
	}

	return ar.GetArticle(ctx, nil, slug, userId)
}

func (ar *articlesRepository) DeleteArticle(ctx context.Context, articleId uuid.UUID) error {
	const sql = `
DELETE FROM articles
WHERE id = UUID_TO_BIN(?)`

	if _, err := ar.db.ExecContext(ctx, sql, articleId); err != nil {
		return err
	}

	return nil
}

func (ar *articlesRepository) GetArticles(ctx context.Context, userId uuid.UUID, tag, author, favorited string, limit, offset int) ([]ArticleCompositeQuery, error) {
	var articles []ArticleCompositeQuery
	{
		// Note: this monster/abomination of a query encapsulates all the data we need to conform
		// for article queries including article information, follower and favorite info, and author info
		const sql = `
SELECT a.id                                                                AS "id",
       a.created_at                                                        AS "created_at",
       a.updated_at                                                        AS "updated_at",
       a.title                                                             AS "title",
       a.body                                                              AS "body",
       a.description                                                       AS "description",
       a.slug                                                              AS "slug",
       u.id                                                                AS "user_id",
       EXISTS(SELECT 1
              FROM user_favorites af
              WHERE af.user_id = UUID_TO_BIN(?)
                AND af.article_id = a.id)                                  AS "favorited",
       (SELECT COUNT(*) FROM user_favorites uf WHERE uf.article_id = a.id) AS "favorites",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE followee_id = a.author_id
                AND follower_id = UUID_TO_BIN(?))                          AS "following_author",
       u.username                                                          AS "author_username",
       u.bio                                                               AS "author_bio",
       u.image                                                             AS "author_image",
       u.email                                                             AS "author_email"
FROM articles a
         JOIN users u ON u.id = a.author_id
WHERE (? = '' OR ? = u.username)
  AND (? = '' OR EXISTS(SELECT 1
                        FROM tags t
                                 JOIN article_tags at ON (t.id, a.id) = (at.tag_id, at.article_id)
                        WHERE t.description = ?))
  AND (? = '' or EXISTS(SELECT 1
                        FROM users favoriting_user
                                 JOIN user_favorites f ON favoriting_user.id = f.user_id
                        WHERE favoriting_user.username = ?)
    )
ORDER BY a.created_at DESC
LIMIT ? OFFSET ?`

		// Unlike Postgres, MySQL can't quite handler positional parameterized arguments,
		// so we need to pass the  same arguments multiple times to reuse parameters in a query
		err := ar.db.SelectContext(
			ctx,
			&articles,
			sql,
			userId,
			userId,
			author,
			author,
			tag,
			tag,
			favorited,
			favorited,
			limit,
			offset)
		if err != nil {
			return articles, shared.ErrorWithContext("error while querying for articles", err)
		}
	}

	return articles, nil
}

func (ar *articlesRepository) GetArticle(ctx context.Context, tx *sqlx.Tx, slug string, userId uuid.UUID) (*ArticleCompositeQuery, error) {
	var article ArticleCompositeQuery
	{
		// Note: this monster/abomination of a query encapsulates all the data we need to conform
		// for article queries including article information, follower and favorite info, and author info
		const sql = `
SELECT a.id                                                                     AS "id",
       a.created_at                                                             AS "created_at",
       a.updated_at                                                             AS "updated_at",
       a.title                                                                  AS "title",
       a.body                                                                   AS "body",
       a.description                                                            AS "description",
       a.slug                                                                   AS "slug",
       u.id                                                                     AS "user_id",
       EXISTS(SELECT 1
              FROM user_favorites af
              WHERE (af.user_id, af.article_id) = (UUID_TO_BIN(?), a.id))       AS "favorited",
       (SELECT count(*) FROM user_favorites WHERE article_id = a.id)            AS "favorites",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE (followee_id, follower_id) = (a.author_id, UUID_TO_BIN(?))) AS "following_author",
       u.username                                                               AS "author_username",
       u.bio                                                                    AS "author_bio",
       u.image                                                                  AS "author_image"
FROM articles a
        LEFT JOIN users u ON u.id = a.author_id
WHERE a.slug = ?`

		// Unlike Postgres, MySQL can't quite handler positional parameterized arguments,
		// so we need to pass the  same arguments multiple times to reuse parameters in a query
		var err error
		{
			if tx != nil {
				err = tx.GetContext(
					ctx,
					&article,
					sql,
					userId,
					userId,
					slug)
			} else {
				err = ar.db.GetContext(
					ctx,
					&article,
					sql,
					userId,
					userId,
					slug)
			}
		}

		if err != nil {
			return &article, err
		}
	}

	return &article, nil
}

func (ar *articlesRepository) AddFavorite(ctx context.Context, articleId, userId uuid.UUID) error {
	const sql = `
INSERT INTO user_favorites (id, user_id, article_id)
VALUES (UUID_TO_BIN(UUID(), TRUE), UUID_TO_BIN(?), UUID_TO_BIN(?))`

	if _, err := ar.db.ExecContext(ctx, sql, userId, articleId); err != nil {
		return err
	}

	return nil
}

func (ar *articlesRepository) DeleteFavorite(ctx context.Context, articleId, userId uuid.UUID) error {
	const sql = `
DELETE FROM user_favorites
WHERE (user_id, article_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))`

	if _, err := ar.db.ExecContext(ctx, sql, userId, articleId); err != nil {
		return err
	}

	return nil
}
