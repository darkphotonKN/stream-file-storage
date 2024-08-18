package p2p

// Represents the type of the arbritrary data sent between two nodes in the
// network via each transport
type Message struct {
	Payload []byte
}
