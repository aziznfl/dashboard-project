package tests

import (
	"fmt"
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestAppError(t *testing.T) {
	t.Run("NewError", func(t *testing.T) {
		err := entity.NewError(entity.ErrorCodeBadRequest, "bad request message")
		assert.Equal(t, entity.ErrorCodeBadRequest, err.Code)
		assert.Equal(t, "bad request message", err.Message)
		assert.Equal(t, "bad request message", err.Error())
	})

	t.Run("WrapError", func(t *testing.T) {
		innerErr := fmt.Errorf("inner error")
		err := entity.WrapError(innerErr, entity.ErrorCodeInternal, "internal error message")
		assert.Equal(t, entity.ErrorCodeInternal, err.Code)
		assert.Equal(t, "internal error message", err.Message)
		assert.Equal(t, "internal error message: inner error", err.Error())
	})

	t.Run("ConvenienceConstructors", func(t *testing.T) {
		assert.Equal(t, entity.ErrorCodeNotFound, entity.ErrorNotFound("not found").Code)
		assert.Equal(t, entity.ErrorCodeUnauthorized, entity.ErrorUnauthorized("unauth").Code)
		assert.Equal(t, entity.ErrorCodeInternal, entity.ErrorInternal("internal").Code)
		assert.Equal(t, entity.ErrorCodeBadRequest, entity.ErrorBadRequest("bad").Code)
	})
}
