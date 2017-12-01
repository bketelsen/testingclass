package mocks

import "github.com/bketelsen/testingclass/interfacesmocks/inventory"

type MockOrderManager struct{}

func (o *MockOrderManager) Get(id int) (*inventory.Order, error) {
	panic("not implemented")
}

func (o *MockOrderManager) Create(or inventory.Order) (*inventory.Order, error) {
	panic("not implemented")
}

func (o *MockOrderManager) Cancel(or *inventory.Order) error {
	panic("not implemented")
}
