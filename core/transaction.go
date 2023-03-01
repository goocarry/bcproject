package core

import (
	"fmt"

	"github.com/goocarry/bcproject/crypto"
	"github.com/goocarry/bcproject/types"
)

// Transaction ...
type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature

	// cached version of the tx hash
	hash types.Hash
}

// NewTransaction creates new Transaction.
func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

// Hash ...
func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}

	return tx.hash
}

// Sign ...
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

// Verify ...
func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
