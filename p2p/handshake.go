package p2p

// import "errors"

// returned if local and remote node connection could not be established
// var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error {
	return nil
}
