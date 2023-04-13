package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
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
		SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error)
		UpdateUser(ctx context.Context, id uuid.UUID, username, email, bio, image, password string) (*UserEntity, error)
	}
	usersRepository struct {
		db *sqlx.DB
	}
)

func (u *UserEntity) ToUser(token string) *users.User {
	return &users.User{
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

	// We can use the username/email key here as we index as a unique key in the database on both fields
	const query string = "SELECT * FROM users WHERE (username, email) = (?, ?)"

	if err := r.db.GetContext(ctx, &user, query, username, email); err != nil {
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
