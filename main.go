package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct{
    Index int       `json:"index"` 
    Transactions []Transaction  `json:"transactions"` 
    PrevHash []byte `json:"prevHash"` 
    Hash []byte     `json:"-"` 
    Nonce uint64    `json:"nonce"`
}

type BlockForHash struct {
    Index    int                `json:"index"`
    Transactions []Transaction  `json:"transactions"` 
    PrevHash []byte             `json:"prevHash"`
    Nonce    uint64             `json:"nonce"`
}

type BlockChain struct{
    Chain []Block
    Difficulty int
}

type Transaction struct{
    From Account    `json:"from"` 
    To Account      `json:"to"` 
    Amount float32  `json:"amount"` 
}

type Account struct{
    Id uint32   `json:"id"` 
    Name string `json:"name"` 
}

func GenerateGenesisBlock(b *Block, difficulty int){
    b.Index=0
    b.Transactions= nil
    b.PrevHash=nil
    b.MineBlock(difficulty)
} 

func InitBlockChain() *BlockChain{

    bc := &BlockChain{
        Chain: make([]Block, 0), 
        Difficulty: 1,
    }
    var genesisBlock Block
    GenerateGenesisBlock(&genesisBlock,bc.Difficulty)
    bc.Chain = append(bc.Chain, genesisBlock) 
    return bc
}



func (b *Block)CalculateHash() error {
    hb := &BlockForHash{
        Index: b.Index,
        Transactions: b.Transactions,
        PrevHash: b.PrevHash,
        Nonce: b.Nonce,
    }
    data, err := json.Marshal(hb) 
    if err != nil {
        return err
    }
    hash := sha256.Sum256(data)
    b.Hash = hash[:]
    return nil
}

func (b *Block)MineBlock(difficulty int){
    b.Nonce = 0
    for {
        b.CalculateHash()
        if CheckHash(b.Hash,difficulty){
            break;
        }
        b.Nonce++
    }
}

func CheckHash(hash []byte,difficulty int) bool {
    for i:=0;i<difficulty;i++{
        if hash[i]!= 0x00{
            return false
        }
    }
    return true
}

func (bc *BlockChain)AddBlockToChain(transactions []Transaction) {
    block := Block{
        Index: len(bc.Chain),
        Transactions: transactions,
    }
    if(len(bc.Chain)!=0){
        block.PrevHash = bc.Chain[block.Index-1].Hash
    }


    block.MineBlock(bc.Difficulty)

    bc.Chain = append(bc.Chain, block)
}

func (bc *BlockChain)DebugBlockChain(){
    for _, v := range bc.Chain{
        fmt.Printf("Index Block: %d\n",v.Index)
        for i,t := range v.Transactions{
            fmt.Printf("Transaction[%d]: \n",i)
            fmt.Printf("Amount: %f\tFrom %d\tTo %d\n",t.Amount,t.From.Id,t.To.Id)
        }
        fmt.Printf("Hash: %x\n", v.Hash)
        fmt.Printf("PrevHash: %x\n", v.PrevHash)
        fmt.Printf("Nonce: %d\n", v.Nonce)
        fmt.Printf("\n----------------------\n\n")
    }
}

func (bc *BlockChain)ValidateBlockChain() bool{
    for i :=range bc.Chain{
        if  i!=0 && (!CheckHash(bc.Chain[i].Hash,bc.Difficulty) ||!bytes.Equal(bc.Chain[i].PrevHash,bc.Chain[i-1].Hash)){
            return false 
        }
    }
    return true;
}

func main(){
    var bc *BlockChain = InitBlockChain() 
    ac1 := Account{
        Id: 1,
        Name: "ali",
    }
    ac2 := Account{
        Id: 2,
        Name: "bob",
    }
    t1 := Transaction{
        From: ac1,
        To: ac2,
        Amount: 10,
    }

    t2 := Transaction{
        From: ac2,
        To: ac1,
        Amount: 22,
    }
    t3 := Transaction{
        From: ac1,
        To: ac2,
        Amount: 33,
    }

    t4 := Transaction{
        From: ac2,
        To: ac1,
        Amount: 44,
    }
    ts1 := []Transaction{t1,t2}
    ts2 := []Transaction{t3,t4}
    bc.AddBlockToChain(ts1)
    bc.AddBlockToChain(ts2)
    bc.DebugBlockChain()
    if bc.ValidateBlockChain() {
        fmt.Println("[+]Blockchain correcta")
    }else {
        fmt.Println("[!]Blockchain incorecta")}
}
