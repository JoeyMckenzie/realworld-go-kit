package users

type UserRequest struct {
	Username *string
	Email    *string
	Password *string
}

type AuthenticationRequest struct {
	User *UserRequest
}
