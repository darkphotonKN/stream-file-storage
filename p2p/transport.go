package p2p

// represents a remote node
type Peer interface {
}

// handles the communication between nodes in a network
// this can be UDP, TCP, websockets, http, etc
type Transport interface {
	ListenAndAccept() error
}
