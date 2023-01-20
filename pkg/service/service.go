package service

import (
	mod "L0WB/internal/domain"
	rep "L0WB/pkg/repository"
)

type SvControllerData interface {
	GetOrder(order_uid string) (mod.Order, error)
	GetAllOrders() (map[string]mod.Order, error)
	InsertOrder(ord *mod.Order) error
}

type Service struct {
	SvControllerData
	Cache map[string]mod.Order
}

func NewService(rep *rep.Repository) *Service {
	return &Service{
		SvControllerData: NewSvController(rep.RpController),
		Cache:            make(map[string]mod.Order),
	}
}
