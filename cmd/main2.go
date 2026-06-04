package main

import (
	"fmt"
	"net/http"

	"github.com/tejas/tejascoin/internal/api"
	"github.com/tejas/tejascoin/internal/blockchain"
)

func main() {

	bc := blockchain.NewBlockchain()

	server := api.NewServer(bc)

	http.HandleFunc(
		"/chain",
		server.GetChain,
	)

	http.HandleFunc(
		"/wallet",
		server.CreateWallet,
	)

	http.HandleFunc(
		"/balance/",
		server.GetBalance,
	)

	http.HandleFunc(
		"/faucet",
		server.Faucet,
	)

	http.HandleFunc(
		"/transaction",
		server.CreateTransaction,
	)

	http.HandleFunc(
		"/mine",
		server.Mine,
	)
	fmt.Println("Listening on :8080")

	http.ListenAndServe(
		":8080",
		nil,
	)
}
