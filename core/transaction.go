package core

import (
	"fmt"

	"github.com/goocarry/bcproject/crypto"
)

// Transaction ...
type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

// Sign ...
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

// Verify ...
func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
