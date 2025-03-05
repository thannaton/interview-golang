package orderhdr

import (
	"github.com/gin-gonic/gin"
	"github.com/thannaton/interview-golang/internal/core/services"
)

type OrderHandler interface {
	Get(c *gin.Context)
}

type orderHandler struct {
	Service services.Service
}

func NewOrderHandler(coreService services.Service) OrderHandler {
	return &orderHandler{
		Service: coreService,
	}
}
