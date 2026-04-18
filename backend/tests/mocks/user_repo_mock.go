package mocks

import (
	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockUserRepo is a mock for repository.UserRepository
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
