package blockchain

import (
	"bytes"
	"example/blockchain/utils"
	"fmt"
)

type BlockChain struct {
	Chain     []Block
	Difficulty int
}

func InitBlockChain() *BlockChain {
	bc := &BlockChain{
		Chain:     make([]Block, 0),
		Difficulty: 1,
	}
	var genesisBlock Block
	GenerateGenesisBlock(&genesisBlock, bc.Difficulty)
	bc.Chain = append(bc.Chain, genesisBlock)
	return bc
}

func (bc *BlockChain) AddBlockToChain(transactions []Transaction) {
	block := Block{
		Index:       len(bc.Chain),
		Transactions: transactions,
	}
	if len(bc.Chain) != 0 {
		block.PrevHash = bc.Chain[block.Index-1].Hash
	}

	block.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, block)
}

func (bc *BlockChain) DebugBlockChain() {
	for _, v := range bc.Chain {
		fmt.Printf("Index Block: %d\n", v.Index)
		for i, t := range v.Transactions {
			fmt.Printf("Transaction[%d]: \n", i)
			fmt.Printf("Amount: %f\tFrom %d\tTo %d\n",t.Amount, t.From, t.To)
		}
		fmt.Printf("Hash: %x\n", v.Hash)
		fmt.Printf("PrevHash: %x\n", v.PrevHash)
		fmt.Printf("Nonce: %d\n", v.Nonce)
		fmt.Printf("\n----------------------\n\n")
	}
}

func (bc *BlockChain) ValidateBlockChain() bool {
	for i := range bc.Chain {
		if i != 0 && (!utils.CheckHash(bc.Chain[i].Hash, bc.Difficulty) || !bytes.Equal(bc.Chain[i].PrevHash, bc.Chain[i-1].Hash)) {
			return false
		}
	}
	return true
}
