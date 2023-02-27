package core

import (
	"io"

	"github.com/goocarry/bcproject/types"
)

// Transaction ...
type Transaction struct {
	Data []byte

	From types.Address
}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(r io.Writer) error {
	return nil
}
