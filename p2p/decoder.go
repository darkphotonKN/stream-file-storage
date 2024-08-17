package p2p

import "io"

type Decoder interface {
	// reads bytes and returns any errors encountered
	Decode(io.Reader, any) error
}
