package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/goocarry/bcproject/crypto"
	"github.com/goocarry/bcproject/types"
)

// Header ...
type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Height        uint32
	Timestamp     int64
}

// Block ...
type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// cached version of the header hash
	hash types.Hash
}

// NewBlock ...
func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

// Sign ...
func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

// Verify ...
func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}

// Decode ...
func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

// Encoder ...
func (b *Block) Encoder(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

// Hash ...
func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

// HeaderData ...
func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}
