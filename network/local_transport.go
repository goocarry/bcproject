package network

import (
	"fmt"
	"sync"
)

// LocalTransport implements Transport interface.
type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

// NewLocalTransport returns new LocalTransport implementing Transport.
func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

// Consume is a method for consuming RPC messages from actual transport.
func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

// Connect ...
func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

// SendMessage ...
func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

// Addr ...
func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
