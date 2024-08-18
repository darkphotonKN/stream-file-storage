package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	// reads bytes and returns any errors encountered
	Decode(io.Reader, any) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, v any) error {

	return gob.NewDecoder(reader).Decode(v)
}

func NewDecoder() Decoder {
	return GOBDecoder{}
}
