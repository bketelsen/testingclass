package postgres

import (
	"database/sql"

	"github.com/bketelsen/testingclass/interfacesmocks/inventory"
)

// Compile-time proof of interface implementation
var _ inventory.OrderStorage = (*OrderService)(nil)

type OrderService struct {
	db *sql.DB
}

func NewOrderService(db *sql.DB) inventory.OrderStorage {
	return &OrderService{db: db}
}

func (s *OrderService) Get(id int) (*inventory.Order, error) {
	panic("not implemented")
}

func (s *OrderService) Create(o inventory.Order) (*inventory.Order, error) {
	panic("not implemented")
}

func (s *OrderService) Cancel(o *inventory.Order) error {
	panic("not implemented")
}
