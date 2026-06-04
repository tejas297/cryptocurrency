package blockchain

import "fmt"

func (bc *Blockchain) ValidateTransaction(
	tx Transaction,
) error {
	// 1. Verify cryptographic ownership
	if err := ValidateSignature(tx); err != nil {
		return err
	}

	// SYSTEM can create coins
	if tx.From == "SYSTEM" {
		return nil
	}

	balance := bc.GetBalance(tx.From)

	if balance < tx.Amount {
		return fmt.Errorf(
			"insufficient funds: balance=%f amount=%f",
			balance,
			tx.Amount,
		)
	}

	return nil
}
