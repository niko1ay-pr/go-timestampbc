package blockchain

type Block struct {
	Index        uint64        `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
	Nonce        uint64        `json:"nonce"`
	Difficulty   uint          `json:"difficulty"`
	Version      string        `json:"version"`
}
