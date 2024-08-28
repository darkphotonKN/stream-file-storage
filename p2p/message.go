package p2p

import "net"

// Represents the type of the arbritrary data sent between two nodes in the
// network via each transport
type RPC struct {
	From    net.Addr
	Payload []byte
}
