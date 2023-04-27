package repositories

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type (
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

    TagsRepository interface {
        CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
        GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error)
        GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error)
        GetArticleTags(ctx context.Context, articleId uuid.UUID) ([]string, error)
    }

    tagsRepository struct {
        db *sqlx.DB
    }
)

func NewTagsRepository(db *sqlx.DB) TagsRepository {
    return &tagsRepository{
        db: db,
    }
}

func (t *tagsRepository) GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    var tag TagEntity
    {
        const sql = "SELECT * FROM tags WHERE description = ?"
        var err error
        {
            if tx != nil {
                err = tx.GetContext(ctx, &tag, sql, description)
            } else {
                err = t.db.GetContext(ctx, &tag, sql, description)
            }
        }

        if err != nil {
            return &tag, shared.ErrorWithContext("error occurred while retrieving tags", err)
        }
    }

    return &tag, nil
}

func (t *tagsRepository) GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    var articleTagEntity ArticleTagEntity
    {
        const sql = "SELECT * FROM article_tags WHERE (tag_id, article_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))"
        var err error
        {
            if tx != nil {
                err = tx.GetContext(ctx, &articleTagEntity, sql, tagId, articleId)
            } else {
                err = t.db.GetContext(ctx, &articleTagEntity, sql, tagId, articleId)
            }
        }

        if err != nil {
            return &articleTagEntity, err
        }
    }

    return &articleTagEntity, nil
}

func (t *tagsRepository) CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    const sql = "INSERT IGNORE INTO tags (id, description) VALUES (UUID_TO_BIN(UUID(), true), ?)"

    if tx != nil {
        if _, err := tx.Exec(sql, description); err != nil {
            return &TagEntity{}, shared.ErrorWithContext("error while attempting to insert tag", err)
        }
    } else {
        if _, err := t.db.Exec(sql, description); err != nil {
            return &TagEntity{}, shared.ErrorWithContext("error while attempting to insert tag", err)
        }
    }

    return t.GetTag(ctx, tx, description)
}

func (t *tagsRepository) CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    const sql = `
INSERT IGNORE INTO article_tags (id, article_id, tag_id)
VALUES (UUID_TO_BIN(UUID(), true), UUID_TO_BIN(?), UUID_TO_BIN(?))`

    if _, err := tx.Exec(sql, articleId, tagId); err != nil {
        return &ArticleTagEntity{}, err
    }

    return t.GetArticleTag(ctx, tx, tagId, articleId)
}

func (t *tagsRepository) GetArticleTags(ctx context.Context, articleId uuid.UUID) ([]string, error) {
    var articleTags []string
    {
        const sql = `
SELECT description FROM tags t
JOIN article_tags at
    ON t.id = at.tag_id
WHERE article_id = UUID_TO_BIN(?)
ORDER BY description`

        if err := t.db.SelectContext(ctx, &articleTags, sql, articleId); err != nil {
            return articleTags, shared.ErrorWithContext(fmt.Sprintf("error while querying for article tags for article %s", articleId), err)
        }
    }

    return articleTags, nil
}
