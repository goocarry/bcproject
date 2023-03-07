package network

import (
	"sync"

	"github.com/goocarry/bcproject/core"
	"github.com/goocarry/bcproject/types"
)

// TxPool ...
type TxPool struct {
	all     *TxSortedMap
	pending *TxSortedMap
	// The maxLength of the total pool of transactions.
	// When the pool is full we will prune the oldest transaction.
	maxLength int
}

// NewTxPool ...
func NewTxPool(maxLength int) *TxPool {
	return &TxPool{
		all:       NewTxSortedMap(),
		pending:   NewTxSortedMap(),
		maxLength: maxLength,
	}
}

// Add adds a transaction to the pool, the caller is responsible checking if the
// tx already exists.
func (p *TxPool) Add(tx *core.Transaction) {
	// prune the oldest transaction that is sitting in the pool
	if p.all.Count() == p.maxLength {
		oldest := p.all.First()
		p.all.Remove(oldest.Hash(core.TxHasher{}))
	}

	if !p.all.Contains(tx.Hash(core.TxHasher{})) {
		p.all.Add(tx)
		p.pending.Add(tx)
	}
}

// Contains returns a boolean if transaction alreaady exists by hash.
func (p *TxPool) Contains(hash types.Hash) bool {
	return p.all.Contains(hash)
}

// Pending returns a slice of transactions that are in the pending pool
func (p *TxPool) Pending() []*core.Transaction {
	return p.pending.txx.Data
}

// ClearPending ...
func (p *TxPool) ClearPending() {
	p.pending.Clear()
}

// PendingCount ...
func (p *TxPool) PendingCount() int {
	return p.pending.Count()
}

// TxSortedMap ...
type TxSortedMap struct {
	lock   sync.RWMutex
	lookup map[types.Hash]*core.Transaction
	txx    *types.List[*core.Transaction]
}

// NewTxSortedMap ...
func NewTxSortedMap() *TxSortedMap {
	return &TxSortedMap{
		lookup: make(map[types.Hash]*core.Transaction),
		txx:    types.NewList[*core.Transaction](),
	}
}

// First returns the oldest transaction from lookup.
func (t *TxSortedMap) First() *core.Transaction {
	t.lock.RLock()
	defer t.lock.RUnlock()

	first := t.txx.Get(0)
	return t.lookup[first.Hash(core.TxHasher{})]
}

// Get returns transaction from lookup by it's hash.
func (t *TxSortedMap) Get(h types.Hash) *core.Transaction {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.lookup[h]
}

// Add insert transaction to TxSortedMap.
func (t *TxSortedMap) Add(tx *core.Transaction) {
	hash := tx.Hash(core.TxHasher{})

	t.lock.Lock()
	defer t.lock.Unlock()

	if _, ok := t.lookup[hash]; !ok {
		t.lookup[hash] = tx
		t.txx.Insert(tx)
	}
}

// Remove deletes transaction from TxSortedMap by it's hash.
func (t *TxSortedMap) Remove(h types.Hash) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.txx.Remove(t.lookup[h])
	delete(t.lookup, h)
}

// Count returns total number of txx.
func (t *TxSortedMap) Count() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return len(t.lookup)
}

// Contains ...
func (t *TxSortedMap) Contains(h types.Hash) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	_, ok := t.lookup[h]
	return ok
}

// Clear ...
func (t *TxSortedMap) Clear() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.lookup = make(map[types.Hash]*core.Transaction)
	t.txx.Clear()
}
