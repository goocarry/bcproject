package core

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Blockchain ...
type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

// NewBlockchain ...
func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

// SetValidator ...
func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

// AddBlock ...
func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

// GetHeader ...
func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height < bc.Height() {
		return nil, fmt.Errorf("given height (%d) too high", height)
	}

	return bc.headers[height], nil
}

// HasBlock ...
func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// Height ...
// [0, 1, 2, 3] => 4 len
// [0, 1, 2, 3] => 3 height
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)

	logrus.WithFields(logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(BlockHasher{}),
	}).Info("adding new block")

	return bc.store.Put(b)
}
