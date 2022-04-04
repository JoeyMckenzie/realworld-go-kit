package services

import (
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

func NewMockTokenService() TokenService {
	return &MockTokenService{}
}

func (m *MockTokenService) GetOptionalUserIdFromAuthorizationHeader(authorizationHeader string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTokenService) GetRequiredUserIdFromAuthorizationHeader(authorizationHeader string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTokenService) GenerateUserToken(id int, email string) (string, error) {
	args := m.Called(id, email)
	return args.String(0), args.Error(1)
}
