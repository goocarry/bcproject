package core

import "fmt"

// Validator ...
type Validator interface {
	ValidateBlock(*Block) error
}

// BlockValidator ...
type BlockValidator struct {
	bc *Blockchain
}

// NewBlockValidator ...
func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

// ValidateBlock ...
func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block (%d) with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block (%s) too high", b.Hash(BlockHasher{}))
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("the hash of the previous block (%s) is invalid", b.PrevBlockHash)
	}

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
