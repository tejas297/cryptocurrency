package api

import (
	"encoding/json"
	"net/http"

	"github.com/tejas/tejascoin/internal/blockchain"
	"github.com/tejas/tejascoin/internal/wallet"
)

type Server struct {
	BC      *blockchain.Blockchain
	Wallets map[string]*wallet.Wallet
}

func NewServer(
	bc *blockchain.Blockchain,
) *Server {
	return &Server{
		BC:      bc,
		Wallets: make(map[string]*wallet.Wallet),
	}
}

func (s *Server) GetChain(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		s.BC.Blocks,
	)
}
