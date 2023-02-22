package types

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

// Hash ...
type Hash [32]uint8

// IsZero returns true if Hash is empty and false vice versa.
func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}

	return true
}

// ToSlice ...
func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}

	return b
}

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

// HashFromBytes ...
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

// RandomBytes generates slice of random bytes.
func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

// RandomHash generates random Hash.
func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
