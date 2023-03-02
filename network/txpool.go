package network

import (
	"sort"

	"github.com/goocarry/bcproject/core"
	"github.com/goocarry/bcproject/types"
)

// TxMapSorter ...
type TxMapSorter struct {
	transactions []*core.Transaction
}

// NewTxMapSorter ...
func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	txx := make([]*core.Transaction, len(txMap))

	i := 0
	for _, val := range txMap {
		txx[i] = val
		i++
	}

	s := &TxMapSorter{txx}
	sort.Sort(s)

	return s
}

// Len ...
func (s *TxMapSorter) Len() int { return len(s.transactions) }

// Swap ...
func (s *TxMapSorter) Swap(i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

// Swap ...
func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen() < s.transactions[j].FirstSeen()
}

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

// Transactions ...
func (p *TxPool) Transactions() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
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
