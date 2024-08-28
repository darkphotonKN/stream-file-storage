package p2p

import (
	"encoding/gob"
	"io"
)

// A decoder just needs to be able to decode messages by taking in a io.Reader
// a pointer to a message to update the message after its decoded.
type Decoder interface {
	// reads bytes and returns any errors encountered
	Decode(io.Reader, *RPC) error
}

// -- Gob Decoder --

func NewGobDecoder() Decoder {
	return GOBDecoder{}
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, rpc *RPC) error {

	return gob.NewDecoder(reader).Decode(rpc)
}

// -- Default Decoder --

func NewDefaultDecoder() Decoder {
	return DefaultDecoder{}
}

type DefaultDecoder struct {
}

func (dec DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {
	buf := make([]byte, 1028) // slightly larger than 1K
	// reads data from into the buffer
	// n is the number of bytes read
	n, err := r.Read(buf)

	if err != nil {
		return err
	}

	// retreive only the used part stored in the buffer
	// from the start to "n"
	rpc.Payload = buf[:n]

	return nil
}
