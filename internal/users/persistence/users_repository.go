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
        IsFollowingUser(ctx context.Context, followerUserId, followeeUserId int) bool
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
select *
from public.user_profile_follows
where (follower_user_id, followee_user_id) = ($1::integer, $2::integer)`

    if err := ur.db.GetContext(ctx, &userFollow, sql, followerUserId, followeeUserId); err != nil {
        return nil, err
    }

    return &userFollow, nil
}

func (ur *usersRepository) CreateUserFollow(ctx context.Context, followerUserId, followeeUserId int) (*UserProfileFollowEntity, error) {
    var userFollow UserProfileFollowEntity

    const sql = `
insert into public.user_profile_follows (created_at, follower_user_id, followee_user_id)
values (current_timestamp, $1::integer, $2::integer)
returning *`

    if err := ur.db.GetContext(ctx, &userFollow, sql, followerUserId, followeeUserId); err != nil {
        return nil, err
    }

    return &userFollow, nil
}

func (ur *usersRepository) RemoveUserFollow(ctx context.Context, followerUserId, followeeUserId int) error {
    const sql = `
delete from public.user_profile_follows
where (follower_user_id, followee_user_id) = ($1::integer, $2::integer)`

    if _, err := ur.db.ExecContext(ctx, sql, followerUserId, followeeUserId); err != nil {
        return err
    }

    return nil
}

func (ur *usersRepository) GetUser(ctx context.Context, userId int) (*UserEntity, error) {
    var userEntity UserEntity

    const sql = `
select *
from public.users u
where u.id = $1::integer
`

    if err := ur.db.GetContext(ctx, &userEntity, sql, userId); err != nil {
        return nil, err
    }

    return &userEntity, nil
}

func (ur *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
    var createdUser UserEntity

    const sql = `
insert into public.users (created_at, updated_at, username, email, password)
values (current_timestamp, current_timestamp, $1::varchar, $2::varchar, $3::varchar)
returning *`

    if err := ur.db.GetContext(ctx, &createdUser, sql, username, email, password); err != nil {
        return nil, err
    }

    return &createdUser, nil
}

func (ur *usersRepository) UpdateUser(ctx context.Context, userId int, username string, email string, password string, bio string, image string) (*UserEntity, error) {
    var user UserEntity

    const sql = `
update public.users
set
    username = $1::varchar,
	email = $2::varchar,
    password = $3::varchar,
    bio = $4::varchar,
    image = $5::varchar,
    updated_at = current_timestamp
where id = $6
returning *;
`
    if err := ur.db.GetContext(ctx, &user, sql, username, email, password, bio, image, userId); err != nil {
        return nil, err
    }

    return &user, nil
}

func (ur *usersRepository) FindUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
    const sql = `
select *
from public.users u
where u.username = $1::varchar`

    return ur.handleFindUserQuery(ctx, sql, username)
}

func (ur *usersRepository) FindUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
    const sql = `
select *
from public.users u
where u.email = $1::varchar`

    return ur.handleFindUserQuery(ctx, sql, email)
}

func (ur *usersRepository) FindUserByUsernameOrEmail(ctx context.Context, username string, email string) (*UserEntity, error) {
    var user UserEntity

    const sql = `
select *
from public.users u
where u.username = $1::varchar
OR u.email = $2::varchar`

    if err := ur.db.GetContext(ctx, &user, sql, username, email); err != nil {
        return nil, err
    }

    return &user, nil
}

func (ur *usersRepository) IsFollowingUser(ctx context.Context, followerUserId, followeeUserId int) bool {
    const sql = `
select 1
from user_profile_follows
where (follower_user_id, followee_user_id) = ($1::integer, $2::integer)`

    if result, err := ur.db.QueryxContext(ctx, sql, followerUserId, followeeUserId); err != nil || !result.Next() {
        return false
    }

    return true
}

func (ur *usersRepository) handleFindUserQuery(ctx context.Context, sql string, criteria string) (*UserEntity, error) {
    var user UserEntity

    if err := ur.db.GetContext(ctx, &user, sql, criteria); err != nil {
        return nil, err
    }

    return &user, nil
}
