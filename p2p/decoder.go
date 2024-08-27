package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	// reads bytes and returns any errors encountered
	Decode(io.Reader, *Message) error
}

// -- Gob Decoder --

func NewGobDecoder() Decoder {
	return GOBDecoder{}
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, msg *Message) error {

	return gob.NewDecoder(reader).Decode(msg)
}

// -- Default Decoder --

func NewDefaultDecoder() Decoder {
	return DefaultDecoder{}
}

type DefaultDecoder struct {
}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1028) // slightly larger than 1K
	// reads data from into the buffer
	// n is the number of bytes read
	n, err := r.Read(buf)

	fmt.Printf("n: %d buf: %v", n, buf)

	if err != nil {
		return err
	}

	// retreive only the used part stored in the buffer
	// from the start to "n"
	msg.Payload = buf[:n]

	return nil
}
