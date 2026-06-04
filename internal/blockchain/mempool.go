package blockchain

func (bc *Blockchain) AddTransaction(
	tx Transaction,
) error {

	if err := bc.ValidateTransaction(tx); err != nil {
		return err
	}

	bc.PendingTransactions =
		append(bc.PendingTransactions, tx)

	return nil
}

func (bc *Blockchain) MinePendingTransactions(
	minerAddress string,
) error {

	rewardTx := Transaction{
		From:   "SYSTEM",
		To:     minerAddress,
		Amount: 50,
	}

	transactions := append(
		bc.PendingTransactions,
		rewardTx,
	)

	err := bc.AddBlock(transactions)
	if err != nil {
		return err
	}

	bc.PendingTransactions = []Transaction{}

	return nil
}
