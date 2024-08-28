package main

import (
	"fmt"
	"log"

	"github.com/darkphotonKN/stream-file-storage/p2p"
)

func main() {

	listenAddr := 5555

	dec := p2p.NewDefaultDecoder()

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: fmt.Sprintf(":%d", listenAddr),
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    dec,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error when connecting to tcp server %s", err.Error())
	}

	select {}

}
