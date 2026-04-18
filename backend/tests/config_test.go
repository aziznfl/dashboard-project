package tests

import (
	"os"
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigLoad(t *testing.T) {
	// Set env vars
	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	os.Setenv("DB_SOURCE", "test.db")
	defer os.Unsetenv("DB_SOURCE")

	cfg := config.Load()

	assert.Equal(t, "test", cfg.AppEnv)
	assert.Equal(t, "test.db", cfg.DbSource)
}

func TestConfigDefault(t *testing.T) {
	// Ensure env vars are not set
	os.Unsetenv("APP_ENV")
	os.Unsetenv("DB_SOURCE")

	cfg := config.Load()

	// These match the defaults in config/env.go
	assert.Equal(t, "production", cfg.AppEnv)
	assert.Equal(t, "dashboard.db", cfg.DbSource)
}
