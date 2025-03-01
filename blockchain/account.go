package blockchain


type Account struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
    BlockChain BlockChain
}
