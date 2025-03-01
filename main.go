package main

import (
	"example/blockchain/blockchain"
	"fmt"
)

func main() {
	var bc *blockchain.BlockChain = blockchain.InitBlockChain()

	ac1 := blockchain.Account{Id: 1, Name: "ali"}
	ac2 := blockchain.Account{Id: 2, Name: "bob"}

	t1 := blockchain.Transaction{From: ac1, To: ac2, Amount: 10}
	t2 := blockchain.Transaction{From: ac2, To: ac1, Amount: 22}
	t3 := blockchain.Transaction{From: ac1, To: ac2, Amount: 33}
	t4 := blockchain.Transaction{From: ac2, To: ac1, Amount: 44}

	ts1 := []blockchain.Transaction{t1, t2}
	ts2 := []blockchain.Transaction{t3, t4}

	bc.AddBlockToChain(ts1)
	bc.AddBlockToChain(ts2)

	bc.DebugBlockChain()

	if bc.ValidateBlockChain() {
		fmt.Println("[+] Blockchain correcta")
	} else {
		fmt.Println("[!] Blockchain incorrecta")
	}
}
