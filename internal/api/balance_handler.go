package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type BalanceResponse struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

func (s *Server) GetBalance(
	w http.ResponseWriter,
	r *http.Request,
) {

	address := strings.TrimPrefix(
		r.URL.Path,
		"/balance/",
	)

	balance := s.BC.GetBalance(
		address,
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		BalanceResponse{
			Address: address,
			Balance: balance,
		},
	)
}
