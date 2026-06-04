package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
)

func Sign(
	privateKey *ecdsa.PrivateKey,
	hash []byte,
) ([]byte, error) {

	return ecdsa.SignASN1(
		rand.Reader,
		privateKey,
		hash,
	)
}

func Verify(
	publicKey *ecdsa.PublicKey,
	hash []byte,
	signature []byte,
) bool {

	return ecdsa.VerifyASN1(
		publicKey,
		hash,
		signature,
	)
}
