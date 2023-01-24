package service

import (
	mod "L0WB/internal/domain"
	"L0WB/pkg/cache"
	rep "L0WB/pkg/repository"
)

type SvControllerData interface {
	GetOrder(order_uid string) (mod.Order, error)
	GetAllOrders() (map[string]mod.Order, error)
	InsertOrder(ord *mod.Order) error
}

type Service struct {
	SvControllerData
	Cache *cache.Cache
}

func NewService(rep *rep.Repository) *Service {
	return &Service{
		SvControllerData: NewSvController(rep.RpController),
		Cache:            cache.New(),
	}
}
