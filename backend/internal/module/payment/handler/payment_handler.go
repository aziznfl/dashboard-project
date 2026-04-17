package handler

import (
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	paymentUsecase "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
)

type PaymentHandler struct {
	paymentUC paymentUsecase.PaymentUsecase
}

func NewPaymentHandler(paymentUC paymentUsecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: paymentUC,
	}
}

func (h *PaymentHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, params openapigen.GetDashboardV1PaymentsParams) {
	filter := repository.PaymentFilter{
		ID:     params.Id,
		Status: params.Status,
		Sort:   params.Sort,
	}

	payments, err := h.paymentUC.ListPayments(r.Context(), filter)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	respPayments := make([]openapigen.Payment, len(payments))
	for i, p := range payments {
		respPayments[i] = openapigen.Payment{
			Id:        &p.ID,
			Amount:    &p.Amount,
			Merchant:  &p.Merchant,
			Status:    &p.Status,
			CreatedAt: &p.CreatedAt,
		}
	}

	transport.WriteJSON(w, http.StatusOK, openapigen.PaymentListResponse{Payments: &respPayments})
}
