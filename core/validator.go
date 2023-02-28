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

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
