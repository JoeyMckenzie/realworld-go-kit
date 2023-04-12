package infrastructure

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
)

type (
	TokenService interface {
		GenerateUserToken(id uuid.UUID, email string) (string, error)
	}

	tokenService struct{}
)

func (ts *tokenService) GenerateUserToken(id uuid.UUID, email string) (string, error) {
	// Generate the token, todo: maybe add more claims like nbf or other validatable auth claims for better security as an example
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   id.String(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenSecret := os.Getenv("TOKEN_SECRET")

	return token.SignedString([]byte(tokenSecret))
}

func NewTokenService() TokenService {
	return &tokenService{}
}