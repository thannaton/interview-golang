package orderhdr

import (
	"github.com/gin-gonic/gin"
	ordersvr "github.com/thannaton/interview-golang/internal/core/services/order"
	"github.com/thannaton/interview-golang/internal/pkg/wrapper"
)

type (
	GetOrderRequest struct {
		No                int     `json:"no" binding:"required"`
		PlatformProductId string  `json:"platformProductId" binding:"required"`
		Qty               int     `json:"qty" binding:"required"`
		UnitPrice         float64 `json:"unitPrice" binding:"required"`
		TotalPrice        float64 `json:"totalPrice" binding:"required"`
	}

	GetOrderResponse struct {
		No         int     `json:"no"`
		ProductId  string  `json:"productId"`
		MaterialId string  `json:"materialId,omitempty"`
		ModelId    string  `json:"modelId,omitempty"`
		Qty        int     `json:"qty"`
		UnitPrice  float64 `json:"unitPrice"`
		TotalPrice float64 `json:"totalPrice"`
	}
)

func (svr *orderHandler) Get(c *gin.Context) {
	var (
		req GetOrderRequest
		res []GetOrderResponse
	)

	defer wrapper.WriteResponse(c, &res)

	// validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Errors = append(c.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
		})
		return
	}

	// call service
	output, err := svr.Service.OrderService().Get(c, ordersvr.GetOrderInput{
		No:                req.No,
		PlatformProductId: req.PlatformProductId,
		Qty:               req.Qty,
		UnitPrice:         req.UnitPrice,
		TotalPrice:        req.TotalPrice,
	})
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypePublic,
		})
		return
	}

	res = parseGetOrderOutput(output)
}

func parseGetOrderOutput(output []ordersvr.GetOrderOutput) []GetOrderResponse {
	items := make([]GetOrderResponse, 0)
	for _, val := range output {
		items = append(items, GetOrderResponse{
			No:         val.No,
			ProductId:  val.ProductId,
			MaterialId: val.MaterialId,
			ModelId:    val.ModelId,
			Qty:        val.Qty,
			UnitPrice:  val.UnitPrice,
			TotalPrice: val.TotalPrice,
		})
	}
	return items
}
