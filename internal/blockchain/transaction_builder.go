package blockchain

import (
	"crypto/ecdsa"

	"github.com/tejas/tejascoin/internal/wallet"
)

func CreateSignedTransaction(
	privateKey *ecdsa.PrivateKey,
	from string,
	to string,
	amount float64,
	publicKey []byte,
) (Transaction, error) {

	tx := Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		PublicKey: publicKey,
	}

	hash := TransactionHash(tx)

	signature, err := wallet.Sign(
		privateKey,
		hash[:],
	)
	if err != nil {
		return Transaction{}, err
	}

	tx.Signature = signature

	return tx, nil
}
