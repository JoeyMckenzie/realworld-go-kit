package shared

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strings"
)

type (
	TokenService interface {
		GenerateUserToken(id uuid.UUID, email string) (string, error)
	}

	tokenService struct{}

	TokenContextKey struct {
		UserId uuid.UUID
	}
)

func NewTokenService() TokenService {
	return &tokenService{}
}

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

func GetUserIdFromAuthorizationHeader(authorizationHeader string) (uuid.UUID, bool) {
	// If there's no authorization header, bail out of attempting to parse the token
	if authorizationHeader == "" {
		return uuid.Nil, false
	}

	// Since we're only looking for an optional header, if we can't parse the token just return default
	token, err := parseTokenFromHeader(authorizationHeader)
	if err != nil {
		return uuid.Nil, false
	}

	// Attempt to parse the claims on the token, bailing out of the process if any of the claims are untypable
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuidClaim, ok := claims["sub"].(string)
		// We'd probably want to handle this scenario better in a real world context, we don't want unexpected panics here
		return uuid.MustParse(uuidClaim), ok
	}

	return uuid.Nil, false
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
		return "", ErrInvalidTokenHeader
	}

	return tokenizedParts[1], nil
}

func getParsedToken(token string) (*jwt.Token, error) {
	// Verify the token, bail out of the request if any parsing failures occur
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, tokenIsValid := token.Method.(*jwt.SigningMethodHMAC); !tokenIsValid {
			return nil, ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
}
