package api

import (
	"net/http"

	"database/sql"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/config"
	"github.com/durianpay/fullstack-boilerplate/internal/infrastructure/cache"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ar "github.com/durianpay/fullstack-boilerplate/internal/module/auth/repository"
	au "github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	ph "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	pr "github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	pu "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
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

func InitAPIHandler(db *sql.DB, redisCache cache.CacheRepository, cfg *config.Config) (*APIHandler, error) {
	jwtExpiredDuration, err := time.ParseDuration(cfg.JwtExpired)
	if err != nil {
		return nil, err
	}

	userRepo := ar.NewUserRepo(db)
	authUC := au.NewAuthUsecase(userRepo, cfg.JwtSecret, jwtExpiredDuration)
	authH := ah.NewAuthHandler(authUC)

	paymentRepo := pr.NewPaymentRepo(db)
	paymentUC := pu.NewPaymentUsecase(paymentRepo, redisCache)
	paymentH := ph.NewPaymentHandler(paymentUC)

	systemH := sh.NewSystemHandler(redisCache)

	return &APIHandler{
		Auth:    authH,
		Payment: paymentH,
		System:  systemH,
	}, nil
}
