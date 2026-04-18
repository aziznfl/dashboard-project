package tests

import (
	"testing"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Manual mock for UserRepository
type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestAuthLogin(t *testing.T) {
	repo := new(mockUserRepo)
	secret := []byte("secret")
	ttl := time.Hour
	authUC := usecase.NewAuthUsecase(repo, secret, ttl)

	password := "password123"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	user := &entity.User{
		ID:           "user-1",
		Email:        "test@example.com",
		PasswordHash: string(hash),
		Role:         "admin",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetUserByEmail", "test@example.com").Return(user, nil).Once()

		token, returnedUser, err := authUC.Login("test@example.com", password)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, user.ID, returnedUser.ID)
		repo.AssertExpectations(t)
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		repo.On("GetUserByEmail", "test@example.com").Return(user, nil).Once()

		token, returnedUser, err := authUC.Login("test@example.com", "wrongpassword")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, returnedUser)
		repo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		repo.On("GetUserByEmail", "notfound@example.com").Return(nil, entity.ErrorNotFound("not found")).Once()

		_, returnedUser, err := authUC.Login("notfound@example.com", password)

		assert.Error(t, err)
		assert.True(t, entity.ErrorCodeNotFound == err.(*entity.AppError).Code)
		assert.Nil(t, returnedUser)
		repo.AssertExpectations(t)
	})
}
