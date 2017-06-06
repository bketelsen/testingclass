package transport

import (
	"net"

	"github.com/bketelsen/testingclass/interfacesmocks/inventory"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
