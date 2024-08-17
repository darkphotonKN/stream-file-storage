package main

import (
	"log"

	"github.com/darkphotonKN/stream-file-storage/p2p"
)

func main() {

	tr := p2p.NewTCPTransport(":5555")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error when connecting to tcp server %s", err.Error())
	}

	select {}

}
