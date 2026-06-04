package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Mine(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req struct {
		Miner string `json:"miner"`
	}

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err := s.BC.
		MinePendingTransactions(
			req.Miner,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Write([]byte(
		"block mined",
	))
}
