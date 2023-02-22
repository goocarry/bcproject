package core

import "io"

// Transaction ...
type Transaction struct {
	Data []byte
}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(r io.Writer) error {
	return nil
}
