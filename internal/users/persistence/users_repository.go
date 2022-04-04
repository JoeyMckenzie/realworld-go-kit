package persistence

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	UserEntity struct {
		Id        int       `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		Bio       string    `db:"bio"`
		Image     string    `db:"image"`
	}

	UserProfileFollowEntity struct {
		Id             int       `db:"id"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		FolloweeUserId int       `db:"followee_user_id"`
		FollowerUserId int       `db:"follower_user_id"`
	}

	UsersRepository interface {
		GetUser(ctx context.Context, userId int) (*UserEntity, error)
		CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error)
		UpdateUser(ctx context.Context, userId int, username, email, password, bio, image string) (*UserEntity, error)
		FindUserByUsername(ctx context.Context, username string) (*UserEntity, error)
		FindUserByEmail(ctx context.Context, email string) (*UserEntity, error)
		FindUserByUsernameOrEmail(ctx context.Context, username, email string) (*UserEntity, error)
		GetUserProfileFollowByFollowee(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error)
		CreateUserFollow(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error)
		RemoveUserFollow(ctx context.Context, followerUserId, followeeUserId int) error
	}

	usersRepository struct {
		db *sqlx.DB
	}

	UsersRepositoryMiddleware func(next UsersRepository) UsersRepository
)

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (ur *usersRepository) GetUserProfileFollowByFollowee(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error) {
	var userFollow UserProfileFollowEntity

	const sql = `
SELECT *
FROM public.user_profile_follows
WHERE (follower_user_id, followee_user_id) = ($1::INTEGER, $2::INTEGER)`

	if err := ur.db.GetContext(ctx, &userFollow, sql, followerUserId, followeeUserId); err != nil {
		return nil, err
	}

	return &userFollow, nil
}

func (ur *usersRepository) CreateUserFollow(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error) {
	var userFollow UserProfileFollowEntity

	const sql = `
INSERT INTO public.user_profile_follows (created_at, updated_at, follower_user_id, followee_user_id)
VALUES (current_timestamp, current_timestamp, $1::INTEGER, $2::INTEGER)
RETURNING *`

	if err := ur.db.GetContext(ctx, &userFollow, sql, followerUserId, followeeUserId); err != nil {
		return nil, err
	}

	return &userFollow, nil
}

func (ur *usersRepository) RemoveUserFollow(ctx context.Context, followerUserId, followeeUserId int) error {
	const sql = `
DELETE FROM public.user_profile_follows
WHERE (follower_user_id, followee_user_id) = ($1::INTEGER, $2::INTEGER)`

	_, err := ur.db.ExecContext(ctx, sql, followerUserId, followeeUserId)
	if err != nil {
		return err
	}

	return nil
}

func (ur *usersRepository) GetUser(ctx context.Context, userId int) (*UserEntity, error) {
	var userEntity UserEntity

	const sql = `
SELECT *
FROM public.users u
WHERE u.id = $1::INTEGER
`

	if err := ur.db.GetContext(ctx, &userEntity, sql, userId); err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func (ur *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
	var createdUser UserEntity

	const sql = `
INSERT INTO public.users (created_at, updated_at, username, email, password)
VALUES (current_timestamp, current_timestamp, $1::VARCHAR, $2::VARCHAR, $3::VARCHAR)
RETURNING *`

	if err := ur.db.GetContext(ctx, &createdUser, sql, username, email, password); err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (ur *usersRepository) UpdateUser(ctx context.Context, userId int, username string, email string, password string, bio string, image string) (*UserEntity, error) {
	var user UserEntity

	const sql = `
UPDATE public.users
SET
    username = $1::VARCHAR,
	email = $2::VARCHAR,
    password = $3::VARCHAR,
    bio = $4::VARCHAR,
    image = $5::VARCHAR,
    updated_at = current_timestamp
WHERE id = $6
RETURNING *;
`
	if err := ur.db.GetContext(ctx, &user, sql, username, email, password, bio, image, userId); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepository) FindUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	const sql = `
SELECT *
FROM public.users u
WHERE u.username = $1::VARCHAR`

	return ur.handleFindUserQuery(ctx, sql, username)
}

func (ur *usersRepository) FindUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	const sql = `
SELECT *
FROM public.users u
WHERE u.email = $1::VARCHAR`

	return ur.handleFindUserQuery(ctx, sql, email)
}

func (ur *usersRepository) FindUserByUsernameOrEmail(ctx context.Context, username string, email string) (*UserEntity, error) {
	var user UserEntity

	const sql = `
SELECT *
FROM public.users u
WHERE u.username = $1::VARCHAR
OR u.email = $2::VARCHAR`

	if err := ur.db.GetContext(ctx, &user, sql, username, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepository) handleFindUserQuery(ctx context.Context, sql string, criteria string) (*UserEntity, error) {
	var user UserEntity

	if err := ur.db.GetContext(ctx, &user, sql, criteria); err != nil {
		return nil, err
	}

	return &user, nil
}
