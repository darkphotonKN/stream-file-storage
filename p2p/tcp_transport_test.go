package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4002"

	tcpOpts := TCPTransportOpts{
		ListenAddr: listenAddr,
		ShakeHands: NOPHandshakeFunc,
		Decoder:    DefaultDecoder{},
	}

	tr := NewTCPTransport(
		tcpOpts,
	).(*TCPTransport)

	assert.Equal(t, tr.TCPTransportOpts.ListenAddr, ":4002")

	// check that this is returning nil and not an error so that server is running
	assert.Nil(t, tr.ListenAndAccept())

}
