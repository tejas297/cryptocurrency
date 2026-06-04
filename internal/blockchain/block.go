package blockchain

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"` // for proof of work
}

// Nonce is a number that miners change to find a hash that meets the difficulty criteria(start with a certain number of leading zeros). It is used in the proof of work algorithm to secure the blockchain and prevent spam transactions.
