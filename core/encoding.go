package core

import (
	"io"
)

// Encoder ...
type Encoder[T any] interface {
	Encode(io.Writer, T) error
}

// Decoder ...
type Decoder[T any] interface {
	Decode(io.Reader, T) error
}
