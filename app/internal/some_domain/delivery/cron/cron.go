package cron

import (
	"context"
	"example-svc/internal/some_domain/delivery"
)

type Handler struct {
	uc delivery.SomeDomain
}

func NewHandler(uc delivery.SomeDomain) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) CloseExpiredSession(ctx context.Context) {
	err := h.uc.CloseExpiredSession(ctx)
	if err != nil {
		return
	}
}
