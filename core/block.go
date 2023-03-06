package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"

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

// Bytes ...
func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

// Block ...
type Block struct {
	*Header
	Transactions []*Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// cached version of the header hash
	hash types.Hash
}

// NewBlock ...
func NewBlock(h *Header, txx []*Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

// NewBlockFromPrevHeader ...
func NewBlockFromPrevHeader(prevHeader *Header, txx []*Transaction) (*Block, error) {
	dataHash, err := CalculateDataHash(txx)
	if err != nil {
		return nil, err
	}
	header := &Header{
		Version:       1,
		Height:        prevHeader.Height + 1,
		DataHash:      dataHash,
		PrevBlockHash: BlockHasher{}.Hash(prevHeader),
		Timestamp:     time.Now().UnixNano(),
	}

	return NewBlock(header, txx), nil
}

// AddTransaction ...
func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, tx)
}

// Sign ...
func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Header.Bytes())
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

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	dataHash, err := CalculateDataHash(b.Transactions)
	if err != nil {
		return err
	}
	if dataHash != b.DataHash {
		return fmt.Errorf("block (%s) has an invalid data hash", b.Hash(BlockHasher{}))
	}

	return nil
}

// Decode ...
func (b *Block) Decode(dec Decoder[*Block]) error {
	return dec.Decode(b)
}

// Encoder ...
func (b *Block) Encoder(enc Encoder[*Block]) error {
	return enc.Encode(b)
}

// Hash ...
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}

// CalculateDataHash ...
func CalculateDataHash(txx []*Transaction) (hash types.Hash, err error) {
	buf := &bytes.Buffer{}

	for _, tx := range txx {
		if err = tx.Encode(NewGobTxEncoder(buf)); err != nil {
			return
		}
	}

	hash = sha256.Sum256(buf.Bytes())

	return
}
