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
	ListPayments(ctx context.Context, filter repository.PaymentFilter) ([]*entity.Payment, *PaginationMeta, error)
}

type PaginationMeta struct {
	Total      int64
	Limit      int
	Page       int
	TotalPages int
	LastID     *string
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

type CacheResult struct {
	Payments []*entity.Payment
	Total    int64
}

func (u *Payment) ListPayments(ctx context.Context, filter repository.PaymentFilter) ([]*entity.Payment, *PaginationMeta, error) {
	// Build cache key based on filter
	id := "all"
	if filter.ID != nil && *filter.ID != "" {
		id = *filter.ID
	}
	status := "all"
	if filter.Status != nil && *filter.Status != "" {
		status = *filter.Status
	}
	merchant := "all"
	if filter.Merchant != nil && *filter.Merchant != "" {
		merchant = *filter.Merchant
	}
	amount := "all"
	if filter.Amount != nil {
		amount = fmt.Sprintf("%d", *filter.Amount)
	}
	sort := "default"
	if filter.Sort != nil && *filter.Sort != "" {
		sort = *filter.Sort
	}
	lastID := "none"
	if filter.LastID != nil && *filter.LastID != "" {
		lastID = *filter.LastID
	}
	cacheKey := fmt.Sprintf("payments:list:%s:%s:%s:%s:%s:%s:%d:%d", id, status, merchant, amount, sort, lastID, filter.Page, filter.Limit)

	var cached CacheResult
	err := u.cache.Get(ctx, cacheKey, &cached)
	if err == nil {
		return cached.Payments, u.buildMeta(cached.Total, filter.Limit, filter.Page, cached.Payments), nil
	}

	total, err := u.repo.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	payments, err := u.repo.List(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	// Cache for 5 minutes
	_ = u.cache.Set(ctx, cacheKey, CacheResult{Payments: payments, Total: total}, 5*time.Minute)
	
	return payments, u.buildMeta(total, filter.Limit, filter.Page, payments), nil
}

func (u *Payment) buildMeta(total int64, limit int, page int, payments []*entity.Payment) *PaginationMeta {
	var lastID *string
	if len(payments) > 0 {
		lastID = &payments[len(payments)-1].ID
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	if page <= 0 {
		page = 1
	}

	return &PaginationMeta{
		Total:      total,
		Limit:      limit,
		Page:       page,
		TotalPages: totalPages,
		LastID:     lastID,
	}
}
