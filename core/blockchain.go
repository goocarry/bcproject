package core

import (
	"fmt"
	"sync"

	"github.com/go-kit/log"
)

// Blockchain ...
type Blockchain struct {
	logger    log.Logger
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

// NewBlockchain ...
func NewBlockchain(l log.Logger, genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
		logger:  l,
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

	bc.lock.Lock()
	defer bc.lock.Unlock()

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
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()
	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()

	bc.logger.Log(
		"msg", "new block",
		"hash", b.Hash(BlockHasher{}),
		"height", b.Height,
		"transactions", len(b.Transactions),
	)

	return bc.store.Put(b)
}
