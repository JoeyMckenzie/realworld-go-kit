package repositories

import (
    "context"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "time"
)

type (
    ArticleEntity struct {
        ID          uuid.UUID
        AuthorID    uuid.UUID `db:"author_id"`
        Slug        string
        Title       string
        Description string
        Body        string
        CreatedAt   string `db:"created_at"`
        UpdatedAt   string `db:"updated_at"`
    }

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

    TagEntity struct {
        ID          uuid.UUID
        Description string
        CreatedAt   string `db:"created_at"`
    }

    ArticleTagEntity struct {
        ID        uuid.UUID
        ArticleID uuid.UUID `db:"article_id"`
        TagID     uuid.UUID `db:"tag_id"`
        CreatedAt string    `db:"created_at"`
    }

    ArticlesRepository interface {
        BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
        GetArticleBySlug(ctx context.Context, tx *sqlx.Tx, slug string) (*ArticleEntity, error)
        CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleEntity, error)
        GetArticles(ctx context.Context, userId uuid.UUID, tag, author, favorited string, limit, offset int) ([]ArticleCompositeQuery, error)
        CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
        GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
        GetArticleTags(ctx context.Context, articleId uuid.UUID) ([]string, error)
    }

    articlesRepository struct {
        db *sqlx.DB
    }
)

func (a ArticleEntity) ToArticle(user *domain.Profile, tagList []string, favoritesCount int, following, favorited bool) (*domain.Article, error) {
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
        Favorited:      favorited,
        FavoritesCount: favoritesCount,
        Author: &domain.Profile{
            Username:  user.Username,
            Email:     user.Email,
            Image:     user.Image,
            Bio:       user.Bio,
            Following: following,
        },
    }, nil
}

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
            Email:     a.Email,
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

func (ar *articlesRepository) GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    var tag TagEntity
    {
        const sql = "SELECT * FROM tags WHERE description = ?"
        var err error
        {
            if tx != nil {
                err = tx.GetContext(ctx, &tag, sql, description)
            } else {
                err = ar.db.GetContext(ctx, &tag, sql, description)
            }
        }

        if err != nil {
            return &tag, shared.ErrorWithContext("error occurred while retrieving tags", err)
        }
    }

    return &tag, nil
}

func (ar *articlesRepository) GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    var articleTagEntity ArticleTagEntity
    {
        const sql = "SELECT * FROM article_tags WHERE (tag_id, article_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))"
        var err error
        {
            if tx != nil {
                err = tx.GetContext(ctx, &articleTagEntity, sql, tagId, articleId)
            } else {
                err = ar.db.GetContext(ctx, &articleTagEntity, sql, tagId, articleId)
            }
        }

        if err != nil {
            return &articleTagEntity, err
        }
    }

    return &articleTagEntity, nil
}

func (ar *articlesRepository) GetArticleBySlug(ctx context.Context, tx *sqlx.Tx, slug string) (*ArticleEntity, error) {
    var articleEntity ArticleEntity
    {
        const sql = "SELECT * FROM articles WHERE slug = ?"

        var err error
        {
            if tx != nil {
                err = tx.GetContext(ctx, &articleEntity, sql, slug)
            } else {
                err = ar.db.GetContext(ctx, &articleEntity, sql, slug)
            }
        }

        if err != nil {
            return &articleEntity, err
        }
    }

    return &articleEntity, nil
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleEntity, error) {
    const sql = `
INSERT INTO articles (id, author_id, slug, title, description, body)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), ?, ?, ?, ?)`

    if _, err := tx.ExecContext(ctx, sql, authorId, slug, title, description, body); err != nil {
        return &ArticleEntity{}, err
    }

    return ar.GetArticleBySlug(ctx, tx, slug)
}

func (ar *articlesRepository) GetArticles(ctx context.Context, userId uuid.UUID, tag, author, favorited string, limit, offset int) ([]ArticleCompositeQuery, error) {
    var articles []ArticleCompositeQuery
    {
        // Note: this monster/abomination of a query encapsulates all the data we need to conform
        // for article queries including article information, follower and favorite info, and author info
        const sql = `
SELECT a.id                                                          AS "id",
       a.created_at                                                  AS "created_at",
       a.updated_at                                                  AS "updated_at",
       a.title                                                       AS "title",
       a.body                                                        AS "body",
       a.description                                                 AS "description",
       a.slug                                                        AS "slug",
       u.id                                                          AS "user_id",
       EXISTS(SELECT 1
              FROM user_favorites af
              WHERE af.user_id = UUID_TO_BIN(?)
                AND af.article_id = a.id)                            AS "favorited",
       (SELECT count(*) FROM user_favorites WHERE article_id = a.id) AS "favorites",
       EXISTS(SELECT 1
              FROM user_follows
              WHERE followee_id = a.author_id
                AND follower_id = UUID_TO_BIN(?))                    AS "following_author",
       u.username                                                    AS "author_username",
       u.bio                                                         AS "author_bio",
       u.image                                                       AS "author_image",
       u.email                                                       AS "author_email"
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

func (ar *articlesRepository) CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    const sql = "INSERT IGNORE INTO tags (id, description) VALUES (UUID_TO_BIN(UUID(), true), ?)"

    if tx != nil {
        if _, err := tx.Exec(sql, description); err != nil {
            return &TagEntity{}, shared.ErrorWithContext("error while attempting to insert tag", err)
        }
    } else {
        if _, err := ar.db.Exec(sql, description); err != nil {
            return &TagEntity{}, shared.ErrorWithContext("error while attempting to insert tag", err)
        }
    }

    return ar.GetTag(ctx, tx, description)
}

func (ar *articlesRepository) CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    const sql = `
INSERT IGNORE INTO article_tags (id, article_id, tag_id)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), UUID_TO_BIN(?))`

    if _, err := tx.Exec(sql, articleId, tagId); err != nil {
        return &ArticleTagEntity{}, err
    }

    return ar.GetArticleTag(ctx, tx, tagId, articleId)
}

func (ar *articlesRepository) GetArticleTags(ctx context.Context, articleId uuid.UUID) ([]string, error) {
    var articleTags []string
    {
        const sql = `
SELECT description FROM tags t
JOIN article_tags at
    ON t.id = at.tag_id
WHERE article_id = UUID_TO_BIN(?)
ORDER BY description`

        if err := ar.db.SelectContext(ctx, &articleTags, sql, articleId); err != nil {
            return articleTags, shared.ErrorWithContext(fmt.Sprintf("error while querying for article tags for article %s", articleId), err)
        }
    }

    return articleTags, nil
}
