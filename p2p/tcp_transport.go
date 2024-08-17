package p2p

import (
	// "bytes"
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
		conn:     conn,
		outbound: outbound,
	}
}

// TCP Transport Struct Options
type TCPTransportOpts struct {
	ListenAddr string
	ShakeHands HandshakeFunc
	Decoder    Decoder
}

// create a TCP transport container
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu sync.RWMutex

	// -- peers --
	// key is network address, which includes:
	// string "Network" to represent network type ("TCP" / "UDP")
	// string "String" to represent the unique ip address
	// value is the Peer which represents the remote node
	peers map[net.Addr]Peer
}

// conforms to the Transport Interface, using DIP for decoupling
func NewTCPTransport(tcpTransportOpts TCPTransportOpts) Transport {
	return &TCPTransport{
		TCPTransportOpts: tcpTransportOpts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.TCPTransportOpts.ListenAddr)

	if err != nil {
		return err
	}

	// go routine to listen for networks
	go t.startAcceptLoop()

	return nil
}

// listens and serves each tcp connection
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		// current connection fails, iterate to start listening again
		if err != nil {
			fmt.Println("Error while attepting to accept connection.", err)

			// break out of current loop
			continue
		}

		// start individual go-routine to handle specific connections
		go t.handleConn(conn)
	}
}

// TODO: update to final message type
type TempMsg struct{}

// serves message within individual TCP connection
func (t *TCPTransport) handleConn(conn net.Conn) {
	// create new tcp connection, outbound peer (making a connection with another peer)
	peer := NewTCPPeer(conn, true)

	fmt.Printf("New incoming connection, peer: %+v\n", peer)

	// attempt to establish handshake
	if err := t.TCPTransportOpts.ShakeHands(conn); err != nil {
		// close connection if handeshake failed

		fmt.Printf("TCP handshake error %s\n", err)
		conn.Close()
		return
	}

	// message read loop - reading from connection
	// buf := new(bytes.Buffer)

	tempMsg := TempMsg{}

	for {
		if err := t.TCPTransportOpts.Decoder.Decode(conn, &tempMsg); err != nil {
			fmt.Printf("Error when decoding incoming message to TCP server %s", err)
			continue
		}

	}

}
