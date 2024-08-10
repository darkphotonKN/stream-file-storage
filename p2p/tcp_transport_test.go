package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4002"

	tr := NewTCPTransport(":4002").(*TCPTransport)

	assert.Equal(t, listenAddr, tr.listenAddress)
}
