package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
)

func ParsePublicKey(
	pubBytes []byte,
) (*ecdsa.PublicKey, error) {

	x, y := elliptic.Unmarshal(
		elliptic.P256(),
		pubBytes,
	)

	if x == nil || y == nil {
		return nil, errors.New(
			"invalid public key",
		)
	}

	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}
