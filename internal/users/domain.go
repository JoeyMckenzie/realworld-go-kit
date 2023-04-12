package users

type (
    User struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Token    string `json:"token"`
        Image    string `json:"image"`
        Bio      string `json:"bio"`
    }

    LoginUserRequest struct {
        Email    *string `validate:"required,email"`
        Password *string `validate:"required"`
    }

    RegisterUserRequest struct {
        Username *string `validate:"required"`
        Email    *string `validate:"required,email"`
        Password *string `validate:"required"`
    }

    UpdateUserRequest struct {
        Username *string
        Email    *string
        Password *string
        Image    *string
        Bio      *string
    }

    AuthenticationRequest[T LoginUserRequest | RegisterUserRequest | UpdateUserRequest] struct {
        User *T `json:"user" validate:"required"`
    }

    AuthenticationResponse struct {
        User *User `json:"user"`
    }
)
