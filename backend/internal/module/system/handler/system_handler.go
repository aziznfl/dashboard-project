package handler

import (
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/infrastructure/cache"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
)

type SystemHandler struct {
	cache cache.CacheRepository
}

func NewSystemHandler(cache cache.CacheRepository) *SystemHandler {
	return &SystemHandler{
		cache: cache,
	}
}

func (h *SystemHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	transport.WriteJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func (h *SystemHandler) PostClearCache(w http.ResponseWriter, r *http.Request) {
	err := h.cache.Flush(r.Context())
	if err != nil {
		transport.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to clear cache"})
		return
	}

	transport.WriteJSON(w, http.StatusOK, map[string]string{"message": "cache cleared successfully"})
}
