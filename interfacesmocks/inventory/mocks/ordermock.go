package mocks

import "github.com/bketelsen/beyondthebasics/src/inventory"

type MockOrderManager struct{}

func (o *MockOrderManager) Get(id int) (*inventory.Order, error) {
	panic("not implemented")
}

func (o *MockOrderManager) Create(o inventory.Order) (*inventory.Order, error) {
	panic("not implemented")
}

func (o *MockOrderManager) Cancel(o *inventory.Order) error {
	panic("not implemented")
}
