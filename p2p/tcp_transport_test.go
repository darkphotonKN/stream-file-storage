package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4002"

	tcpOpts := TCPTransportOpts{
		ListenAddr: ":4002",
		ShakeHands: NOPHandshakeFunc,
	}

	tr := NewTCPTransport(
		tcpOpts,
	).(*TCPTransport)

	assert.Equal(t, listenAddr, tr.TCPTransportOpts.ListenAddr)

	// check that this is returning nil and not an error
	assert.Nil(t, tr.ListenAndAccept())
}
