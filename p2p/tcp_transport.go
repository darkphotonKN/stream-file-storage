package p2p

import (
	"fmt"
	"net"
	"sync"
)

/**
* TCP Transport Protocol
**/

// TCP Peer
type TCPPeer struct {
	// underlying connection for peer
	conn net.Conn

	// represents the type of responsiblity of dialing connection as a node
	// if we dial and receive a connection = outbound then this value is true
	// if we dial and receiev a connection = inbound then this value is false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}

}

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

// conforms to the Transport Interface, using DIP for decoupling
func NewTCPTransport(listenAddr string) Transport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	// go routine to listen for networks
	go t.startAcceptLoop()

	return nil
}

// listens and serves the tcp connection
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Println("Error while attepting to accept connection.", err)

			// break out of current loop
			continue
		}

		// start another go-routine to handle specific connections
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("Connection established %v\n", conn)
}
