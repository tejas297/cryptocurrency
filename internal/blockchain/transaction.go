package blockchain

type Transaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	PublicKey []byte  `json:"publicKey"`
	Signature []byte  `json:"signature"`
}
