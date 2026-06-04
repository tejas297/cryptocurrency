package wallet

import (
	"crypto/sha256"
	"encoding/hex"
)

func AddressFromPublicKey(
	pub []byte,
) string {

	hash := sha256.Sum256(pub)

	return hex.EncodeToString(
		hash[:],
	)
}
