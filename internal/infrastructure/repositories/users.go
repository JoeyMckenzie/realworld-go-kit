package repositories

import (
    "context"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
)

type (
    UserEntity struct {
        ID        uuid.UUID
        Username  string
        Email     string
        Password  string
        Image     string
        Bio       string
        CreatedAt string `db:"created_at"`
        UpdatedAt string `db:"updated_at"`
    }

    UsersRepository interface {
        CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error)
        GetUserById(ctx context.Context, id uuid.UUID) (*UserEntity, error)
        GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*UserEntity, error)
        GetUserByUsername(ctx context.Context, username string) (*UserEntity, error)
        SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error)
        UpdateUser(ctx context.Context, id uuid.UUID, username, email, bio, image, password string) (*UserEntity, error)
        GetExistingFollow(ctx context.Context, followerId, followeeId uuid.UUID) (uuid.UUID, error)
        AddFollow(ctx context.Context, followerId, followeeId uuid.UUID) error
        DeleteFollow(ctx context.Context, followerId, followeeId uuid.UUID) error
    }
    usersRepository struct {
        db *sqlx.DB
    }
)

func (u *UserEntity) ToUser(token string) *domain.User {
    return &domain.User{
        ID:       u.ID,
        Username: u.Username,
        Email:    u.Email,
        Token:    token,
        Image:    u.Image,
        Bio:      u.Bio,
    }
}

func (u *UserEntity) ToProfile(following bool) *domain.Profile {
    return &domain.Profile{
        Username:  u.Username,
        Image:     u.Image,
        Bio:       u.Bio,
        Following: following,
    }
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
    return &usersRepository{
        db: db,
    }
}

func (r *usersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
    var user []UserEntity
    const query string = "SELECT * FROM users WHERE username = ? OR email = ?"

    if err := r.db.SelectContext(ctx, &user, query, username, email); err != nil {
        return user, err
    }

    return user, nil
}

func (r *usersRepository) GetUserById(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
    var user UserEntity
    const query string = "SELECT * FROM users WHERE id = UUID_TO_BIN(?) LIMIT 1"

    if err := r.db.GetContext(ctx, &user, query, id); err != nil {
        return &user, err
    }

    return &user, nil
}

func (r *usersRepository) GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*UserEntity, error) {
    var user UserEntity

    // We can use the username/email key here as we index as a unique key in the data on both fields
    const query string = "SELECT * FROM users WHERE (username, email) = (?, ?)"

    if err := r.db.GetContext(ctx, &user, query, username, email); err != nil {
        return &user, err
    }

    return &user, nil
}

func (r *usersRepository) GetUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
    var user UserEntity
    const query string = "SELECT * FROM users WHERE username = ?"

    if err := r.db.GetContext(ctx, &user, query, username); err != nil {
        return &user, err
    }

    return &user, nil
}

func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
    const command string = "INSERT INTO users (id, username, email, password) VALUES (UUID_TO_BIN(UUID()), ?, ?, ?)"

    if _, err := r.db.ExecContext(ctx, command, username, email, password); err != nil {
        return &UserEntity{}, err
    }

    return r.GetUserByUsernameAndEmail(ctx, username, email)
}

func (r *usersRepository) UpdateUser(ctx context.Context, id uuid.UUID, username, email, bio, image, password string) (*UserEntity, error) {
    const command string = `
UPDATE users
SET username = ?,
    email = ?,
    bio = ?,
	image = ?,
	password = ?,
	updated_at = CURRENT_TIMESTAMP
WHERE id = UUID_TO_BIN(?)`

    if _, err := r.db.ExecContext(ctx, command, username, email, bio, image, password, id); err != nil {
        return &UserEntity{}, nil
    }

    return r.GetUserById(ctx, id)
}

func (r *usersRepository) GetExistingFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) (uuid.UUID, error) {
    var followId uuid.UUID
    const query string = "SELECT id FROM user_follows WHERE (follower_id, followee_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))"

    if err := r.db.GetContext(ctx, &followId, query, followerId, followeeId); err != nil {
        return uuid.Nil, err
    }

    return followId, nil
}

func (r *usersRepository) AddFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
    const command string = `
INSERT INTO user_follows (id, follower_id, followee_id)
VALUES (UUID_TO_BIN(UUID()), UUID_TO_BIN(?), UUID_TO_BIN(?))`

    if _, err := r.db.ExecContext(ctx, command, followerId, followeeId); err != nil {
        return err
    }

    return nil
}

func (r *usersRepository) DeleteFollow(ctx context.Context, followerId uuid.UUID, followeeId uuid.UUID) error {
    const command string = `
DELETE FROM user_follows
WHERE (follower_id, followee_id) = (UUID_TO_BIN(?), UUID_TO_BIN(?))`

    if _, err := r.db.ExecContext(ctx, command, followerId, followeeId); err != nil {
        return err
    }

    return nil
}
