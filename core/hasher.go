package core

import (
	"crypto/sha256"

	"github.com/goocarry/bcproject/types"
)

// Hasher ...
type Hasher[T any] interface {
	Hash(T) types.Hash
}

// BlockHasher ...
type BlockHasher struct{}

// Hash ...
func (BlockHasher) Hash(b *Header) types.Hash {
	h := sha256.Sum256(b.Bytes())
	return types.Hash(h)
}
