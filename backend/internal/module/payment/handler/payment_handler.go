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
	limit := 10

	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	filter := repository.PaymentFilter{
		ID:       params.Id,
		Status:   params.Status,
		Merchant: params.Merchant,
		Amount:   intPtrToInt64Ptr(params.Amount),
		Sort:     params.Sort,
		LastID:   params.LastId,
		Page:     page,
		Limit:    limit,
	}

	payments, meta, err := h.paymentUC.ListPayments(r.Context(), filter)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	respPayments := make([]openapigen.Payment, len(payments))
	for i, p := range payments {
		amount := int(p.Amount)
		respPayments[i] = openapigen.Payment{
			Id:        &p.ID,
			Amount:    &amount,
			Merchant:  &p.Merchant,
			Status:    &p.Status,
			CreatedAt: &p.CreatedAt,
		}
	}

	total := int(meta.Total)
	transport.WriteJSON(w, http.StatusOK, openapigen.PaymentListResponse{
		Data: &respPayments,
		Meta: &openapigen.PaginationMeta{
			Total:      total,
			Limit:      meta.Limit,
			Page:       meta.Page,
			TotalPages: meta.TotalPages,
			LastId:     meta.LastID,
		},
	})
}

func intPtrToInt64Ptr(v *int) *int64 {
	if v == nil {
		return nil
	}
	res := int64(*v)
	return &res
}
