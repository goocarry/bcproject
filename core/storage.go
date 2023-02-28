package core

// Storage ...
type Storage interface {
	Put(*Block) error
}

// MemoryStore ...
type MemoryStore struct {
}

// NewMemoryStore ...
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// Put ...
func (ms MemoryStore) Put(b *Block) error {
	return nil
}
