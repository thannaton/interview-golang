package ordersvr

import "github.com/gin-gonic/gin"

type OrderService interface {
	Get(c *gin.Context, input GetOrderInput) ([]GetOrderOutput, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}
