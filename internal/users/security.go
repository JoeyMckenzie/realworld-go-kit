package users

import "golang.org/x/crypto/bcrypt"

type (
	// SecurityService handles the hashing and verification of passwords
	//when authenticating users during login or registration.
	SecurityService interface {
		HashPassword(rawPassword string) (string, error)
		IsValidPassword(existingPassword, rawPassword string) bool
	}

	securityService struct{}
)

func NewSecurityService() SecurityService {
	return &securityService{}
}

func (s *securityService) HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *securityService) IsValidPassword(existingPassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(rawPassword))
	return err == nil
}
