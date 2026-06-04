package blockchain

import (
	"strings"
)

func MineBlock(block Block, difficulty int) Block {

	target := strings.Repeat("0", difficulty)

	for {

		hash := CalculateHash(block)

		if strings.HasPrefix(hash, target) {
			block.Hash = hash
			return block
		}

		block.Nonce++
	}
}
