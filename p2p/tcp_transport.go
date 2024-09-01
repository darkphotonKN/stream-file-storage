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

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	// underlying connection for peer
	conn net.Conn

	// represents the type of responsiblity of dialing connection as a node
	// if we dial and receive a connection = outbound then this value is true
	// if we dial and receiev a connection = inbound then this value is false
	outbound bool
}

// for closing peer node in the network
func (p *TCPPeer) Close() error {
	// attempt to close the connection to the node, returning the error if it fails
	return p.conn.Close()
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
	// provide callback access to the Peer object
	OnPeer func(Peer) error
}

// create a TCP transport container
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC

	mu sync.RWMutex

	// -- peers --
	// key is network address, which includes:
	// string "Network" to represent network type ("TCP" / "UDP")
	// string "String" to represent the unique ip address
	// value is the Peer which represents the remote node
	// NOTE: moving this to let server handle list of peers instead (for now)
	// peers map[net.Addr]Peer
}

// conforms to the Transport Interface, using DIP for decoupling
func NewTCPTransport(opts TCPTransportOpts) Transport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Restricts caller to only listen through the return type of this
// [read-only] channel, for reading the messages received from another peer
// in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	// creates a tcp listener that is capable of listening on a tcp connection
	t.listener, err = net.Listen("tcp", t.TCPTransportOpts.ListenAddr)

	if err != nil {
		return err
	}

	// goroutine to listen for networks,
	// used to allow other incoming requests while my goroutine
	// handles more incoming request for this tcp connection
	go t.startAcceptLoop()

	return nil
}

// listens and serves each tcp connection
func (t *TCPTransport) startAcceptLoop() {
	for {
		// actually use the tcp listener to start and listen to incoming CONNECTIONS
		conn, err := t.listener.Accept()

		// current connection fails, iterate to start listening again
		if err != nil {
			fmt.Println("Error while attepting to accept connection.", err)

			// break out of current loop
			continue
		}

		// start individual go-routine to handle individual CONNECTION'S requests
		go t.handleConn(conn)
	}
}

// serves message within individual TCP connection
func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("Dropping the peer connection, reason: %s", err)
		conn.Close()
	}()

	// create new tcp connection, outbound peer (making a connection with another peer)
	peer := NewTCPPeer(conn, true)

	fmt.Printf("[TCP Server Msg] New incoming connection, peer: %v\n\n", peer)

	// Only establish read loop after handshake and onPeer existance is confirmed

	// attempt to establish handshake
	if err = t.TCPTransportOpts.ShakeHands(conn); err != nil {

		// close connection if handeshake failed

		fmt.Printf("TCP handshake error %s\n", err)
		conn.Close()
		return
	}

	// check onPeer when provided does not have an error for a legit peer connection
	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// message read loop - reading from connection
	// buf := new(bytes.Buffer)

	rpc := RPC{}

	for {
		if err = t.TCPTransportOpts.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("Error when decoding incoming message to TCP server %s", err)
			continue
		}

		fmt.Printf("[TCP Server Msg] Message recieved in tcp connection %+v\n", rpc)
	}

}
