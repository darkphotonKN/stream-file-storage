package main

import (
	"fmt"
	"log"

	"github.com/darkphotonKN/stream-file-storage/p2p"
)

func onPeer(p2p.Peer) error {
	// do some logic with the peer outside of transport
	return nil
}

func main() {
	listenAddr := 5555

	dec := p2p.NewDefaultDecoder()

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: fmt.Sprintf(":%d", listenAddr),
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    dec,
		OnPeer:     onPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	// NOTE: Remove after testing
	go func() {
		// listen to the TCP server's responses

		serverMsg := <-tr.Consume()

		fmt.Printf("Channel message received: %s", serverMsg)

	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error when connecting to tcp server %s", err.Error())
	}

	select {}
}
