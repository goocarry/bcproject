package util

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/goocarry/bcproject/core"
	"github.com/goocarry/bcproject/crypto"
	"github.com/goocarry/bcproject/types"
	"github.com/stretchr/testify/assert"
)

// RandomBytes generates slice of random bytes.
func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

// RandomHash generates random Hash.
func RandomHash() types.Hash {
	return types.HashFromBytes(RandomBytes(32))
}

// NewRandomTransaction return a new random transaction whithout signature.
func NewRandomTransaction(size int) *core.Transaction {
	return core.NewTransaction(RandomBytes(size))
}

// NewRandomTransactionWithSignature ...
func NewRandomTransactionWithSignature(t *testing.T, privKey crypto.PrivateKey, size int) *core.Transaction {
	tx := NewRandomTransaction(size)
	assert.Nil(t, tx.Sign(privKey))
	return tx
}

// NewRandomBlock ...
func NewRandomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *core.Block {
	txSigner := crypto.GeneratePrivateKey()
	tx := NewRandomTransactionWithSignature(t, txSigner, 100)
	header := &core.Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	b := core.NewBlock(header, []*core.Transaction{tx})
	dataHash, err := core.CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash

	return b
}

// NewRandomBlockWithSignature ...
func NewRandomBlockWithSignature(t *testing.T, pk crypto.PrivateKey, height uint32, prevHash types.Hash) *core.Block {
	b := NewRandomBlock(t, height, prevHash)
	assert.Nil(t, b.Sign(pk))

	return b
}
