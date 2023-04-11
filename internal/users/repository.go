package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
		GetUser(ctx context.Context, id uuid.UUID) (*UserEntity, error)
		SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error)
	}
)

func (u *UserEntity) ToUser(token string) *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
		Image:    u.Image,
		Bio:      u.Bio,
	}
}

func NewRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

type usersRepository struct {
	db *sqlx.DB
}

func (r *usersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
	var user []UserEntity
	const query string = "SELECT * FROM users WHERE username = ? OR email = ?"

	if err := r.db.SelectContext(ctx, &user, query, username, email); err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersRepository) GetUser(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	var user *UserEntity
	const query string = "SELECT * FROM users WHERE id = ? LIMIT 1"

	if err := r.db.SelectContext(ctx, &user, query, id); err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
	var user UserEntity
	const command string = "INSERT INTO users (users.id, users.username, users.email, users.password) VALUES (UUID_TO_BIN(UUID()), ?, ?, ?)"
	const query string = `
SELECT id,
       username,
       email,
       password,
       image,
       bio,
       created_at,
       updated_at
FROM users
WHERE id = LAST_INSERT_ID()
LIMIT 1`

	if _, err := r.db.ExecContext(ctx, command, username, email, password); err != nil {
		return &user, err
	}

	if err := r.db.GetContext(ctx, &user, query); err != nil {
		return &user, nil
	}

	return &user, nil
}
