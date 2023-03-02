package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/goocarry/bcproject/core"
)

// MessageType ...
type MessageType byte

const (
	// MessageTypeTx ...
	MessageTypeTx MessageType = 0x1
	// MessageTypeBock ...
	MessageTypeBock
)

// RPC is a message sent over the transport layer.
type RPC struct {
	From    NetAddr
	Payload io.Reader
}

// Message ...
type Message struct {
	Header MessageType
	Data   []byte
}

// NewMessage ...
func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

// Bytes ...
func (msg *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

// RPCHandler ...
type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

// DefaultRPCHandler ...
type DefaultRPCHandler struct {
	p RPCProcessor
}

// NewDefaultRPCHandler ...
func NewDefaultRPCHandler(p RPCProcessor) *DefaultRPCHandler {
	return &DefaultRPCHandler{
		p: p,
	}
}

// HandleRPC ...
func (h *DefaultRPCHandler) HandleRPC(rpc RPC) error {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return err
	}

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return fmt.Errorf("failed to decode message from %s: %s", rpc.From, err)
		}
		return h.p.ProcessTransaction(rpc.From, tx)
	default:
		return fmt.Errorf("invalid message header %x", msg.Header)
	}
}

// RPCProcessor ...
type RPCProcessor interface {
	ProcessTransaction(NetAddr, *core.Transaction) error
}
