package repository

import (
	con "L0WB/internal/controller"
	mod "L0WB/internal/domain"
	"github.com/jmoiron/sqlx"
)

type RpController interface {
	GetOrder(order_uid string) (mod.Order, error)
	GetAllOrders() (map[string]mod.Order, error)
	InsertOrder(ord *mod.Order) error
}

type Repository struct {
	RpController
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RpController: con.NewDbController(db),
	}
}
