package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/goocarry/bcproject/core"
	"github.com/sirupsen/logrus"
)

// MessageType ...
type MessageType byte

const (
	// MessageTypeTx ...
	MessageTypeTx MessageType = 0x1
	// MessageTypeBlock ...
	MessageTypeBlock MessageType = 0x2
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

// DecodedMessage ...
type DecodedMessage struct {
	From NetAddr
	Data any
}

// RPCDecodeFunc ...
type RPCDecodeFunc func(RPC) (*DecodedMessage, error)

// DefaRPCDecodeFunc ...
func DefaRPCDecodeFunc(rpc RPC) (*DecodedMessage, error) {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to decode message from %s: %s", rpc.From, err)
	}

	logrus.WithFields(logrus.Fields{
		"from": rpc.From,
		"type": msg.Header,
	}).Debug("new incoming message")

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}

		return &DecodedMessage{
			From: rpc.From,
			Data: tx,
		}, nil

	case MessageTypeBlock:
		block := new(core.Block)
		if err := block.Decode(core.NewGobBlockDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}
		return &DecodedMessage{
			From: rpc.From,
			Data: block,
		}, nil

	default:
		return nil, fmt.Errorf("invalid message header %x", msg.Header)
	}
}

// RPCProcessor ...
type RPCProcessor interface {
	ProcessMessage(*DecodedMessage) error
}
