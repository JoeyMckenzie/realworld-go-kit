package users

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockTokenService struct {
	mock.Mock
}

func (m *mockTokenService) ResetMocks() {
	m.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (m *mockTokenService) GenerateUserToken(id uuid.UUID, email string) (string, error) {
	args := m.Called(id, email)
	return args.String(0), args.Error(1)
}
