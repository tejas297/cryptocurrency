package api

import (
	"encoding/json"
	"net/http"

	"github.com/tejas/tejascoin/internal/wallet"
)

type WalletResponse struct {
	Address string `json:"address"`
}

func (s *Server) CreateWallet(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	newWallet, err := wallet.NewWallet()
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	s.Wallets[newWallet.Address] =
		newWallet

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		WalletResponse{
			Address: newWallet.Address,
		},
	)
}
