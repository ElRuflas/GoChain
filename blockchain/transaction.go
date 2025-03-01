package blockchain

type Transaction struct {
	From   uint32   `json:"from"`
	To     uint32   `json:"to"`
	Amount float32  `json:"amount"`
}
