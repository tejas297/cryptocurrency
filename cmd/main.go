package main

import (
	"fmt"

	"github.com/tejas/tejascoin/internal/wallet"
)

func main1() {
	//create a new blockchain
	// bc := blockchain.NewBlockchain()

	// fmt.Println("Blockchain created")

	// fmt.Printf("%+v\n", bc.Blocks[0])

	// bc := blockchain.NewBlockchain()

	// fmt.Println("Blockchain created")

	// fmt.Printf("%+v\n", bc.Blocks[0])

	// tx := blockchain.Transaction{
	// 	From:   "Alice",
	// 	To:     "Bob",
	// 	Amount: 10,
	// }

	// bc.AddBlock([]blockchain.Transaction{tx})

	// for _, block := range bc.Blocks {
	// 	fmt.Printf("%+v\n\n", block)
	// }

	alice, _ := wallet.NewWallet()

	bob, _ := wallet.NewWallet()

	fmt.Println("Alice")
	fmt.Println(alice.Address)

	fmt.Println()

	fmt.Println("Bob")
	fmt.Println(bob.Address)
}
