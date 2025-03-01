package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct{
    Index int `json:"index"` 
    Data []byte `json:"data"` 
    PrevHash []byte `json:"prevHash"` 
    Hash []byte `json:"-"` 
    Nonce uint64 `json:"nonce"`
}

type BlockChain struct{
    Chain []Block
    Difficulty int
}

func GenerateGenesisBlock(b *Block, difficulty int){
    b.Index=0
    b.Data = []byte("Genesis")
    b.PrevHash=nil
    b.MineBlock(difficulty)
} 

func InitBlockChain() *BlockChain{

    bc := &BlockChain{
        Chain: make([]Block, 0), 
        Difficulty: 3,
    }
    var genesisBlock Block
    GenerateGenesisBlock(&genesisBlock,bc.Difficulty)
    bc.Chain = append(bc.Chain, genesisBlock) 
    return bc
}



func (b *Block)CalculateHash() error {
    data, err := json.Marshal(b) 
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
    hexHash := fmt.Sprintf("%x", hash)
    for i:=0;i<difficulty;i++{
        if hexHash[i]!= '0'{
            return false
        }
    }
    return true
}

func (bc *BlockChain)AddBlockChain(data []byte) {
    block := Block{
        Index: len(bc.Chain),
        Data: data,
    }
    if(len(bc.Chain)!=0){
        block.PrevHash = bc.Chain[block.Index-1].Hash
    }


    block.MineBlock(bc.Difficulty)

    bc.Chain = append(bc.Chain, block)
}

func (bc *BlockChain)DebugBlockChain(){
    for _, v := range bc.Chain{
        fmt.Printf("Index: %d\n",v.Index)
        fmt.Printf("Data: %s\n", v.Data) 
        fmt.Printf("Hash: %x\n", v.Hash)
        fmt.Printf("PrevHash: %x\n", v.PrevHash)
        fmt.Printf("Nonce: %d\n", v.Nonce)
        fmt.Printf("\n----------------------\n\n")
    }
}

func (bc *BlockChain)ValidateBlockChain() bool{
    for i :=range bc.Chain{
        if !CheckHash(bc.Chain[i].Hash,bc.Difficulty) || i!=0 && !bytes.Equal(bc.Chain[i].PrevHash,bc.Chain[i-1].Hash){
            return false 
        }
    }
    return true;
}

func main(){
    var bc *BlockChain = InitBlockChain() 
    bc.AddBlockChain([]byte("Primero"))
    //bc.AddBlockChain([]byte("segundo"))
    bc.DebugBlockChain()
    if bc.ValidateBlockChain() {
        fmt.Println("[+]Blockchain correcta")
    }else {
        fmt.Println("[!]Blockchain incorecta")}
}
