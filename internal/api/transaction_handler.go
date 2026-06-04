package api

import (
	"encoding/json"
	"net/http"

	"github.com/tejas/tejascoin/internal/blockchain"
)

type TransactionRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (s *Server) CreateTransaction(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req TransactionRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	sender, ok := s.Wallets[req.From]
	if !ok {

		http.Error(
			w,
			"wallet not found",
			http.StatusBadRequest,
		)

		return
	}

	pubBytes, err :=
		sender.PublicKey.Bytes()

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	tx, err :=
		blockchain.CreateSignedTransaction(
			sender.PrivateKey,
			sender.Address,
			req.To,
			req.Amount,
			pubBytes,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err = s.BC.AddTransaction(tx)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Write([]byte(
		"transaction added to mempool",
	))
}
