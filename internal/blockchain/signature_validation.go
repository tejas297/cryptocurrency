package blockchain

import (
	"fmt"

	"github.com/tejas/tejascoin/internal/wallet"
)

func ValidateSignature(
	tx Transaction,
) error {

	if tx.From == "SYSTEM" {
		return nil
	}

	pubKey, err := wallet.ParsePublicKey(
		tx.PublicKey,
	)
	if err != nil {
		return err
	}

	derivedAddress :=
		wallet.AddressFromPublicKey(
			tx.PublicKey,
		)

	if derivedAddress != tx.From {
		return fmt.Errorf(
			"public key does not match sender",
		)
	}

	hash := TransactionHash(tx)

	ok := wallet.Verify(
		pubKey,
		hash[:],
		tx.Signature,
	)

	if !ok {
		return fmt.Errorf(
			"invalid signature",
		)
	}

	return nil
}
