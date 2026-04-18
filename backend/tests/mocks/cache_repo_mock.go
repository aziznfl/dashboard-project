package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockCacheRepo is a mock for cache.CacheRepository
type MockCacheRepo struct {
	mock.Mock
}

func (m *MockCacheRepo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockCacheRepo) Get(ctx context.Context, key string, dest interface{}) error {
	args := m.Called(ctx, key, dest)
	return args.Error(0)
}

func (m *MockCacheRepo) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *MockCacheRepo) Flush(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
