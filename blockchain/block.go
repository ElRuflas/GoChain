package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"example/blockchain/utils"
)

type Block struct {
	Index       int           `json:"index"`
	Transactions []Transaction `json:"transactions"`
	PrevHash    []byte        `json:"prevHash"`
	Hash        []byte        `json:"-"`
	Nonce       uint64        `json:"nonce"`
}

type BlockForHash struct {
	Index       int           `json:"index"`
	Transactions []Transaction `json:"transactions"`
	PrevHash    []byte        `json:"prevHash"`
	Nonce       uint64        `json:"nonce"`
}

func GenerateGenesisBlock(b *Block, difficulty int) {
	b.Index = 0
	b.Transactions = nil
	b.PrevHash = nil
	b.MineBlock(difficulty)
}

func (b *Block) CalculateHash() error {
	hb := &BlockForHash{
		Index:       b.Index,
		Transactions: b.Transactions,
		PrevHash:    b.PrevHash,
		Nonce:       b.Nonce,
	}
	data, err := json.Marshal(hb)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
	return nil
}

func (b *Block) MineBlock(difficulty int) {
	b.Nonce = 0
	for {
		b.CalculateHash()
		if utils.CheckHash(b.Hash, difficulty) {
			break
		}
		b.Nonce++
	}
}



