package cache

import (
	"context"
	"time"
)

// CacheRepository defines the interface for caching operations
type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Delete(ctx context.Context, key string) error
	Flush(ctx context.Context) error
}
