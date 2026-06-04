package main

// this test file should verify :

// ✓ Wallet creation
// ✓ Initial coin creation (SYSTEM)
// ✓ Balance calculation
// ✓ Signed transaction creation
// ✓ Transaction validation
// ✓ Mempool
// ✓ Mining
// ✓ Mining rewards
// ✓ Ownership verification
// ✓ Tamper detection

import (
	"fmt"

	"github.com/tejas/tejascoin/internal/blockchain"
	"github.com/tejas/tejascoin/internal/wallet"
)

func main2() {

	// Create blockchain
	bc := blockchain.NewBlockchain()

	// Create wallets
	alice, _ := wallet.NewWallet()
	bob, _ := wallet.NewWallet()
	miner, _ := wallet.NewWallet()

	fmt.Println("========== WALLETS ==========")
	fmt.Println("Alice :", alice.Address)
	fmt.Println("Bob   :", bob.Address)
	fmt.Println("Miner :", miner.Address)

	// Give Alice 100 coins
	err := bc.AddBlock([]blockchain.Transaction{
		{
			From:   "SYSTEM",
			To:     alice.Address,
			Amount: 100,
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("\n========== INITIAL BALANCES ==========")
	fmt.Println("Alice:", bc.GetBalance(alice.Address))
	fmt.Println("Bob  :", bc.GetBalance(bob.Address))
	fmt.Println("Miner:", bc.GetBalance(miner.Address))

	// Create signed transaction
	pubBytes, err := alice.PublicKey.Bytes()
	if err != nil {
		panic(err)
	}

	tx, err := blockchain.CreateSignedTransaction(
		alice.PrivateKey,
		alice.Address,
		bob.Address,
		25,
		pubBytes,
	)

	if err != nil {
		panic(err)
	}

	// Add to mempool
	err = bc.AddTransaction(tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n========== MEMPOOL ==========")
	fmt.Println("Pending transactions:", len(bc.PendingTransactions))

	// Mine block
	err = bc.MinePendingTransactions(
		miner.Address,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("\n========== AFTER MINING ==========")
	fmt.Println("Alice:", bc.GetBalance(alice.Address))
	fmt.Println("Bob  :", bc.GetBalance(bob.Address))
	fmt.Println("Miner:", bc.GetBalance(miner.Address))

	fmt.Println("\n========== BLOCKCHAIN ==========")

	for _, block := range bc.Blocks {

		fmt.Println("--------------------------------")

		fmt.Printf(
			"Index=%d\nHash=%s\nPrevHash=%s\nNonce=%d\nTxCount=%d\n",
			block.Index,
			block.Hash,
			block.PrevHash,
			block.Nonce,
			len(block.Transactions),
		)
	}
}
