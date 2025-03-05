package handler

import "github.com/gin-gonic/gin"

func NewRouter(router *gin.RouterGroup, handler Handler) {
	router.POST("/order", handler.OrderHandler().Get)
}
