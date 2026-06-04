package api

import (
	"encoding/json"
	"net/http"

	"github.com/tejas/tejascoin/internal/blockchain"
)

type FaucetRequest struct {
	Address string `json:"address"`
}

func (s *Server) Faucet(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req FaucetRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err := s.BC.AddBlock(
		[]blockchain.Transaction{
			{
				From:   "SYSTEM",
				To:     req.Address,
				Amount: 100,
			},
		},
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.Write([]byte("100 TJC granted"))
}
