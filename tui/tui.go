package tui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"example/blockchain/blockchain" // Ajusta la importación a tu proyecto
)

// Tui es la estructura que contiene los componentes de la TUI
type Tui struct {
	app  *tview.Application
	view *tview.TextView
}

// NewTui crea una nueva instancia de Tui
func NewTui() *Tui {
	return &Tui{
		app:  tview.NewApplication(),
		view: tview.NewTextView(),
	}
}

// ShowBlockchain muestra la blockchain en la interfaz de usuario
func (t *Tui) ShowBlockchain(bc *blockchain.BlockChain) {
	// Configurar el TextView para mostrar los bloques con desplazamiento
	t.view.
		SetDynamicColors(true).
		SetScrollable(true).
		SetTextAlign(tview.AlignLeft).
		SetBackgroundColor(tcell.ColorBlack)

	// Construir el texto de la blockchain
	var blockchainText string
	for _, block := range bc.Chain {
		blockchainText += fmt.Sprintf("[white]Block %d\n", block.Index)
		blockchainText += fmt.Sprintf("[green]PrevHash: [gray]%x\n", block.PrevHash)
		blockchainText += fmt.Sprintf("[blue]Hash: [gray]%x\n", block.Hash)
		blockchainText += fmt.Sprintf("[red]Nonce: [gray]%d\n", block.Nonce)

		if len(block.Transactions) > 0 {
			blockchainText += "[yellow]Transactions:\n"
			for i, tx := range block.Transactions {
				blockchainText += fmt.Sprintf(
					"  [white]%d. From: %d → To: %d | Amount: %.2f\n",
					i+1, tx.From, tx.To, tx.Amount,
				)
			}
		}
		blockchainText += "\n"
	}

	// Asignar el texto al TextView
	t.view.SetText(blockchainText)
}

// CreateTui inicia la interfaz de usuario
func (t *Tui) CreateTui(bc *blockchain.BlockChain) error {
	t.ShowBlockchain(bc)
	t.app.SetRoot(t.view, true)

	// Ejecutar la aplicación
	return t.app.Run()
}

// Stop detiene la aplicación
func (t *Tui) Stop() {
	t.app.Stop()
}
