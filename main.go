package main

import (
	"example/blockchain/blockchain"
    "example/blockchain/tui"
	"fmt"
)

func main() {
	var bc *blockchain.BlockChain = blockchain.InitBlockChain()

	ac1 := blockchain.Account{Id: 1}
	ac2 := blockchain.Account{Id: 2}

	t1 := blockchain.Transaction{From: ac1.Id, To: ac2.Id, Amount: 11}
	t2 := blockchain.Transaction{From: ac2.Id, To: ac1.Id, Amount: 22}
	t3 := blockchain.Transaction{From: ac1.Id, To: ac2.Id, Amount: 33}
	t4 := blockchain.Transaction{From: ac2.Id, To: ac1.Id, Amount: 44}

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

    // Inicializar la TUI
    tuiApp := tui.NewTui()

    // Crear la interfaz de usuario
    if err := tuiApp.CreateTui(bc); err != nil {
        fmt.Println("Error al iniciar la TUI:", err)
    }
}
