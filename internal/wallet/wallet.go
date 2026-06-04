package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

func NewWallet() (*Wallet, error) {

	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(publicKeyBytes)

	address := hex.EncodeToString(hash[:])

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
		Address:    address,
	}, nil
}
