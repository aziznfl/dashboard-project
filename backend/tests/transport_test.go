package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"message": "success"}
	status := http.StatusOK

	transport.WriteJSON(w, status, data)

	assert.Equal(t, status, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp["message"])
}

func TestCodeToStatus(t *testing.T) {
	tests := []struct {
		code     entity.Code
		expected int
	}{
		{entity.ErrorCodeBadRequest, http.StatusBadRequest},
		{entity.ErrorCodeUnauthorized, http.StatusUnauthorized},
		{entity.ErrorCodeInternal, http.StatusInternalServerError},
		{entity.Code("unknown"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(string(tt.code), func(t *testing.T) {
			assert.Equal(t, tt.expected, transport.CodeToStatus(tt.code))
		})
	}
}

func TestWriteError(t *testing.T) {
	t.Run("AppError", func(t *testing.T) {
		w := httptest.NewRecorder()
		appErr := entity.ErrorBadRequest("invalid input")

		transport.WriteError(w, appErr)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp transport.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, string(entity.ErrorCodeBadRequest), resp.Code)
		assert.Equal(t, "invalid input", resp.Message)
	})

	t.Run("GenericError", func(t *testing.T) {
		w := httptest.NewRecorder()
		genErr := assert.AnError

		transport.WriteError(w, genErr)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp transport.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, string(entity.ErrorCodeInternal), resp.Code)
	})
}
