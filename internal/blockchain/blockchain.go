package blockchain

import "time"

type Blockchain struct {
	Blocks              []Block
	PendingTransactions []Transaction
}

// NewBlockchain initializes a new blockchain with a genesis block.
// very block chain start with block 0. which is called genesis.
func NewBlockchain() *Blockchain {

	genesis := Block{
		Index:     0,
		Timestamp: "GENESIS",
		PrevHash:  "",
		Nonce:     0,
	}

	genesis.Hash = CalculateHash(genesis)

	return &Blockchain{
		Blocks:              []Block{genesis},
		PendingTransactions: []Transaction{},
	}
}

func (bc *Blockchain) AddBlock(
	transactions []Transaction,
) error {

	prev := bc.Blocks[len(bc.Blocks)-1]

	block := Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prev.Hash,
	}

	for _, tx := range transactions {

		if err := bc.ValidateTransaction(tx); err != nil {
			return err
		}
	}

	block = MineBlock(block, 4)

	bc.Blocks = append(
		bc.Blocks,
		block,
	)
	return nil
}
