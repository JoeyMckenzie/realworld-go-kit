package utilities

import "errors"

var (
	ErrMock                      = errors.New("stub err")
	ErrEmailEmpty                = errors.New("email cannot be empty")
	ErrEmailInvalid              = errors.New("email is invalid")
	ErrUsernameNil               = errors.New("username must be provided")
	ErrUsernameEmpty             = errors.New("username cannot be empty")
	ErrPasswordEmpty             = errors.New("password cannot be empty")
	ErrInvalidLoginAttempt       = errors.New("email or password is invalid")
	ErrUserNotFound              = errors.New("user was not found")
	ErrArticlesNotFound          = errors.New("no articles were found")
	ErrCommentNotFound           = errors.New("no comments were found")
	ErrUsernameOrEmailTaken      = errors.New("username or email already exists")
	ErrArticleTitleExists        = errors.New("article with that title already exists")
	ErrInvalidSigningMethod      = errors.New("invalid token signing method")
	ErrUnauthorized              = errors.New("unauthorized attempt to access resource")
	ErrUsernameNotProvided       = errors.New("username was not provided on the request")
	ErrInternalServerError       = errors.New("unexpected error occurred")
	ErrInvalidRequestBody        = errors.New("request body is invalid")
	ErrInvalidLimitOrOffsetValue = errors.New("limit or offset value is not a valid integer")
	ErrInvalidTokenHeader        = errors.New("authorization token is malformed")
	ErrCannotFollowSelf          = errors.New("cannot follow or unfollow self")
	ErrNilInput                  = errors.New("cannot pass nil value")
	ErrForbiddenArticleUpdate    = errors.New("forbidden to update this article")
)
