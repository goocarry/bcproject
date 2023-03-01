package network

import (
	"github.com/goocarry/bcproject/core"
	"github.com/goocarry/bcproject/types"
)

// TxPool ...
type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

// NewTxPool ...
func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction, 0),
	}
}

// Add adds a transaction to the pool, the caller is responsible checking if the
// tx already exists.
func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	p.transactions[hash] = tx

	return nil
}

// Has returns a boolean if transaction alreaady exists by hash.
func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}

// Len ...
func (p *TxPool) Len() int {
	return len(p.transactions)
}

// Flush ...
func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}
