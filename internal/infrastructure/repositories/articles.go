package repositories

import (
    "context"
    "errors"
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
        CommitTransaction(tx *sqlx.Tx) error
        RollbackTransaction(tx *sqlx.Tx) error
        GetArticleBySlug(ctx context.Context, tx *sqlx.Tx, slug string) (*ArticleEntity, error)
        CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleEntity, error)
        CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
        GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
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

func NewArticlesRepository(db *sqlx.DB) ArticlesRepository {
    return &articlesRepository{
        db: db,
    }
}

func (ar *articlesRepository) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
    return ar.db.BeginTxx(ctx, nil)
}

func (ar *articlesRepository) CommitTransaction(tx *sqlx.Tx) error {
    return tx.Commit()
}

func (ar *articlesRepository) RollbackTransaction(tx *sqlx.Tx) error {
    return tx.Rollback()
}

func (ar *articlesRepository) GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    var tag TagEntity
    const sql string = "SELECT * FROM tags WHERE description = ?"

    var err error
    {
        if tx != nil {
            err = tx.GetContext(ctx, &tag, sql, description)
        } else {
            err = ar.db.GetContext(ctx, &tag, sql, description)
        }
    }

    if err != nil {
        return &tag, err
    }

    return &tag, nil
}

func (ar *articlesRepository) GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    var articleTagEntity ArticleTagEntity
    const sql string = "SELECT * FROM article_tags WHERE (tag_id, article_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))"

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

    return &articleTagEntity, nil
}

func (ar *articlesRepository) GetArticleBySlug(ctx context.Context, tx *sqlx.Tx, slug string) (*ArticleEntity, error) {
    var articleEntity ArticleEntity
    const sql string = "SELECT * FROM articles WHERE slug = ?"

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

    return &articleEntity, nil
}

func (ar *articlesRepository) CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleEntity, error) {
    const sql string = `
INSERT INTO articles (id, author_id, slug, title, description, body)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), ?, ?, ?, ?)`

    if _, err := tx.ExecContext(ctx, sql, authorId, slug, title, description, body); err != nil {
        return &ArticleEntity{}, err
    }

    return ar.GetArticleBySlug(ctx, tx, slug)
}

func (ar *articlesRepository) CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    const sql string = "INSERT IGNORE INTO tags (id, description) VALUES (UUID_TO_BIN(UUID(), true), ?)"

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
    const sql string = `
INSERT IGNORE INTO article_tags (id, article_id, tag_id)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), UUID_TO_BIN(?))`

    if _, err := tx.Exec(sql, articleId, tagId); err != nil {
        return &ArticleTagEntity{}, err
    }

    return ar.GetArticleTag(ctx, tx, tagId, articleId)
}
