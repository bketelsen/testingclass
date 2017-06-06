package acme

import "github.com/bketelsen/testingclass/interfacesmocks/inventory"

// Compile-time proof of interface implementation
var _ inventory.SupplierService = (*AcmeClientService)(nil)

type AcmeClientService struct {
	URL string
}

func NewClient(url string) inventory.SupplierService {
	return &AcmeClientService{URL: url}
}

func (a *AcmeClientService) PlaceOrder(o *inventory.Order) error {
	panic("not implemented")
}

func (a *AcmeClientService) GetStatus(o *inventory.Order) error {
	panic("not implemented")
}
