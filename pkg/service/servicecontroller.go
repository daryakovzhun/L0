package service

import (
	mod "L0WB/internal/domain"
	rep "L0WB/pkg/repository"
)

type SvController struct {
	rep rep.RpController
}

func NewSvController(rep rep.RpController) *SvController {
	return &SvController{rep: rep}
}

func (SC *SvController) GetOrder(order_uid string) (mod.Order, error) {
	ord, err := SC.rep.GetOrder(order_uid)
	if err != nil {
		return mod.Order{}, err
	}
	return ord, err
}

func (SC *SvController) InsertOrder(ord *mod.Order) error {
	err := SC.rep.InsertOrder(ord)
	if err != nil {
		return err
	}
	return nil
}

func (SC *SvController) GetAllOrders() (map[string]mod.Order, error) {
	ord, err := SC.rep.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return ord, err
}
