package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

// Encoder ...
type Encoder[T any] interface {
	Encode(T) error
}

// Decoder ...
type Decoder[T any] interface {
	Decode(T) error
}

// GobTxEncoder ...
type GobTxEncoder struct {
	w io.Writer
}

// NewGobTxEncoder ...
func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	gob.Register(elliptic.P256())
	return &GobTxEncoder{
		w: w,
	}
}

// Encode ...
func (e *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(e.w).Encode(tx)
}

// GobTxDecoder ...
type GobTxDecoder struct {
	r io.Reader
}

// NewGobTxDecoder ...
func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	gob.Register(elliptic.P256())
	return &GobTxDecoder{
		r: r,
	}
}

// Decode ...
func (d *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(d.r).Decode(tx)
}
