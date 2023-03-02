package network

// NetAddr ...
type NetAddr string

// Transport is a general interface for transport layer.
type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
