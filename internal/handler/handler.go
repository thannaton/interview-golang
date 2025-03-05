package handler

import (
	"github.com/thannaton/interview-golang/internal/core/services"
	orderhdr "github.com/thannaton/interview-golang/internal/handler/order"
)

type Handler interface {
	OrderHandler() orderhdr.OrderHandler
}

type handler struct {
	Service services.Service
}

func NewHandler(coreService services.Service) Handler {
	return &handler{
		Service: coreService,
	}
}

func (h *handler) OrderHandler() orderhdr.OrderHandler {
	return orderhdr.NewOrderHandler(h.Service)
}
