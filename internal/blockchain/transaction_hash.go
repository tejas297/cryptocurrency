package blockchain

import (
	"crypto/sha256"
	"fmt"
)

func TransactionHash(tx Transaction) [32]byte {

	record := fmt.Sprintf(
		"%s%s%f%x",
		tx.From,
		tx.To,
		tx.Amount,
		tx.PublicKey,
	)

	return sha256.Sum256([]byte(record))
}
