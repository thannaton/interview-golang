package services

import ordersvr "github.com/thannaton/interview-golang/internal/core/services/order"

type Service interface {
	OrderService() ordersvr.OrderService
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) OrderService() ordersvr.OrderService {
	return ordersvr.NewOrderService()
}
