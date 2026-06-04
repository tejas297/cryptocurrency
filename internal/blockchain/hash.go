package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// bitcoin use double sha256, but for simplicity we will use single sha256 here.
func CalculateHash(block Block) string {

	record := fmt.Sprintf(
		"%d%s%v%s%d",
		block.Index,
		block.Timestamp,
		block.Transactions,
		block.PrevHash,
		block.Nonce,
	)

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}
