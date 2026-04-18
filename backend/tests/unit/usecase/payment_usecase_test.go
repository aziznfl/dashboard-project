package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPaymentList_Caching(t *testing.T) {
	repo := new(mocks.MockPaymentRepo)
	cache := new(mocks.MockCacheRepo)
	uc := usecase.NewPaymentUsecase(repo, cache)

	ctx := context.Background()
	filter := repository.PaymentFilter{}
	cacheKey := "payments:list:all:all"

	payments := []*entity.Payment{
		{ID: "1", Amount: "100", Merchant: "A", Status: "success"},
		{ID: "2", Amount: "200", Merchant: "B", Status: "pending"},
	}

	t.Run("CacheHit", func(t *testing.T) {
		// Expect Get from cache and return data
		cache.On("Get", ctx, cacheKey, mock.Anything).Run(func(args mock.Arguments) {
			dest := args.Get(2).(*[]*entity.Payment)
			*dest = payments
		}).Return(nil).Once()

		res, err := uc.ListPayments(ctx, filter)

		assert.NoError(t, err)
		assert.Equal(t, payments, res)
		repo.AssertNotCalled(t, "List", mock.Anything, mock.Anything)
		cache.AssertExpectations(t)
	})

	t.Run("CacheMiss", func(t *testing.T) {
		// Reset mocks
		repo.ExpectedCalls = nil
		cache.ExpectedCalls = nil

		// Expect Get (miss), then Repo call, then Set cache
		cache.On("Get", ctx, cacheKey, mock.Anything).Return(assert.AnError).Once()
		repo.On("List", ctx, filter).Return(payments, nil).Once()
		cache.On("Set", ctx, cacheKey, payments, 5*time.Minute).Return(nil).Once()

		res, err := uc.ListPayments(ctx, filter)

		assert.NoError(t, err)
		assert.Equal(t, payments, res)
		repo.AssertExpectations(t)
		cache.AssertExpectations(t)
	})
}
