package mocks

import (
	"context"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	"github.com/stretchr/testify/mock"
)

// MockPaymentRepo is a mock for repository.PaymentRepository
type MockPaymentRepo struct {
	mock.Mock
}

func (m *MockPaymentRepo) List(ctx context.Context, filter repository.PaymentFilter) ([]*entity.Payment, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Payment), args.Error(1)
}
