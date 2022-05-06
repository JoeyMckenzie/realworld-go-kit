package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"os"
	"strconv"
	"strings"
)

type (
	TokenService interface {
		GetOptionalUserIdFromAuthorizationHeader(authorizationHeader string) (int, error)
		GetRequiredUserIdFromAuthorizationHeader(authorizationHeader string) (int, error)
		GenerateUserToken(id int, email string) (string, error)
	}

	tokenService struct{}
)

func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) GetOptionalUserIdFromAuthorizationHeader(authorizationHeader string) (int, error) {
	if authorizationHeader == "" {
		return -1, nil
	}

	return ts.GetRequiredUserIdFromAuthorizationHeader(authorizationHeader)
}

func (ts *tokenService) GetRequiredUserIdFromAuthorizationHeader(authorizationHeader string) (int, error) {
	token, err := parseTokenFromHeader(authorizationHeader)

	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%v", claims["sub"]), 10, 64)

		if err != nil {
			return -1, utilities.ErrInvalidTokenHeader
		}

		return int(userId), nil
	}

	return -1, utilities.ErrInvalidTokenHeader
}

func (ts *tokenService) GenerateUserToken(id int, email string) (string, error) {
	// Generate the token, todo: maybe add more claims like nbf or other validatable auth claims for better security as an example
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenSecret := os.Getenv("TOKEN_SECRET")

	return token.SignedString([]byte(tokenSecret))
}

func parseTokenFromHeader(authorizationTokenHeader string) (*jwt.Token, error) {
	token, err := getTokenFromHeader(authorizationTokenHeader)

	if err != nil {
		return nil, err
	}

	return getParsedToken(token)
}

func getTokenFromHeader(authorizationTokenHeader string) (string, error) {
	tokenizedParts := strings.Split(authorizationTokenHeader, " ")

	if len(tokenizedParts) != 2 {
		return "", utilities.ErrInvalidTokenHeader
	}

	return tokenizedParts[1], nil
}

func getParsedToken(token string) (*jwt.Token, error) {
	// Verify the token, bail out of the request if any parsing utilities occur
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, tokenIsValid := token.Method.(*jwt.SigningMethodHMAC); !tokenIsValid {
			return nil, utilities.ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
}
