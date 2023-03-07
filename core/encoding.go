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
	return &GobTxDecoder{
		r: r,
	}
}

// Decode ...
func (d *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(d.r).Decode(tx)
}

// GobBlockEncoder ...
type GobBlockEncoder struct {
	w io.Writer
}

// NewGobBlockEncoder ...
func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	return &GobBlockEncoder{
		w: w,
	}
}

// Encode ...
func (e *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(e.w).Encode(b)
}

// GobBlockDecoder ...
type GobBlockDecoder struct {
	r io.Reader
}

// NewGobBlockDecoder ...
func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {
	return &GobBlockDecoder{
		r: r,
	}
}

// Decode ...
func (d *GobBlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(d.r).Decode(b)
}

func init() {
	gob.Register(elliptic.P256())
}
