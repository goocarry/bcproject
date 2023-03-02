package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/goocarry/bcproject/types"
)

// PrivateKey ...
type PrivateKey struct {
	key *ecdsa.PrivateKey
}

// Sign ...
func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		R: r,
		S: s,
	}, nil
}

// GeneratePrivateKey ...
func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

// PublicKey ...
func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		Key: &k.key.PublicKey,
	}
}

// PublicKey ...
type PublicKey struct {
	Key *ecdsa.PublicKey
}

// ToSlice ...
func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key, k.Key.X, k.Key.Y)
}

// Address ...
func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

// Signature ...
type Signature struct {
	S, R *big.Int
}

// Verify ...
func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.Key, data, sig.R, sig.S)
}
