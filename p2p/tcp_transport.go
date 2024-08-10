package p2p

import (
	"net"
	"sync"
)

/**
* TCP Transport Protocol
**/

// create a TCP transport container
type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu sync.RWMutex

	// -- peers --
	// key is network address, which includes:
	// string "Network" to represent network type ("TCP" / "UDP")
	// string "String" to represent the unique ip address
	// value is the Peer which represents the remote node
	peers map[net.Addr]Peer
}

// conforms to the Transport Inteface, using DIP for decoupling
func NewTCPTransport(listenAddr string) Transport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}
