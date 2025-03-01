package blockchain

type Transaction struct {
	From   Account  `json:"from"`
	To     Account  `json:"to"`
	Amount float32  `json:"amount"`
}
