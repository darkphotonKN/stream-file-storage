package main

import (
	"io"
	"log"

	"github.com/darkphotonKN/stream-file-storage/p2p"
)

type Decoder struct {
}

func (d *Decoder) Decode(io.Reader, any) error {
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":5555",
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    &Decoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error when connecting to tcp server %s", err.Error())
	}

	select {}

}
