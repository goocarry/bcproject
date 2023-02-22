package network

// NetAddr ...
type NetAddr string

// RPC is a message sent over the transport layer.
type RPC struct {
	From    NetAddr
	Payload []byte
}

// Transport is a general interface for transport layer.
type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
