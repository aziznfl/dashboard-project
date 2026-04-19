package api

import (
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ph "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	sh "github.com/durianpay/fullstack-boilerplate/internal/module/system/handler"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth    *ah.AuthHandler
	Payment *ph.PaymentHandler
	System  *sh.SystemHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostClearCache(w http.ResponseWriter, r *http.Request) {
	h.System.PostClearCache(w, r)
}

func (h *APIHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	h.System.GetHealth(w, r)
}

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, body openapigen.GetDashboardV1PaymentsParams) {
	h.Payment.GetDashboardV1Payments(w, r, body)
}
