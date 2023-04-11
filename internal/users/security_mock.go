package users

import (
	"github.com/stretchr/testify/mock"
)

type mockSecurityService struct {
	mock.Mock
}

func (m *mockSecurityService) ResetMocks() {
	m.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *mockSecurityService) HashPassword(rawPassword string) (string, error) {
	args := m.Called(rawPassword)
	return args.String(0), args.Error(1)
}

func (m *mockSecurityService) IsValidPassword(existingPassword, rawPassword string) bool {
	args := m.Called(existingPassword, rawPassword)
	return args.Bool(0)
}
