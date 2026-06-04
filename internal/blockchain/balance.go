package blockchain

func (bc *Blockchain) GetBalance(
	address string,
) float64 {

	var balance float64

	for _, block := range bc.Blocks {

		for _, tx := range block.Transactions {

			if tx.To == address {
				balance += tx.Amount
			}

			if tx.From == address {
				balance -= tx.Amount
			}
		}
	}

	return balance
}
