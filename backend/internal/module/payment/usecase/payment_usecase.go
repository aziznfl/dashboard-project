package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/infrastructure/cache"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type PaymentUsecase interface {
	ListPayments(ctx context.Context, filter repository.PaymentFilter) ([]*entity.Payment, error)
}

type Payment struct {
	repo  repository.PaymentRepository
	cache cache.CacheRepository
}

func NewPaymentUsecase(repo repository.PaymentRepository, cache cache.CacheRepository) *Payment {
	return &Payment{
		repo:  repo,
		cache: cache,
	}
}

func (u *Payment) ListPayments(ctx context.Context, filter repository.PaymentFilter) ([]*entity.Payment, error) {
	// Build cache key based on filter
	id := "all"
	if filter.ID != nil {
		id = *filter.ID
	}
	status := "all"
	if filter.Status != nil {
		status = *filter.Status
	}
	cacheKey := fmt.Sprintf("payments:list:%s:%s", id, status)

	var payments []*entity.Payment
	err := u.cache.Get(ctx, cacheKey, &payments)
	if err == nil {
		return payments, nil
	}

	payments, err = u.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Cache for 5 minutes
	_ = u.cache.Set(ctx, cacheKey, payments, 5*time.Minute)

	return payments, nil
}
