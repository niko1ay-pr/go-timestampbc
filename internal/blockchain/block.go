package blockchain

const (
	version          = "1.0"
	hashAlgorithm    = "sha256"
	genesisTimestamp = 12345678
)

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

func NewBlock(index uint64, timestamp int64, prevHash string, difficulty uint) *Block {
	return &Block{
		Index:      index,
		Timestamp:  timestamp,
		PrevHash:   prevHash,
		Difficulty: difficulty,
		Version:    version,
	}
}

func NewGenesisBlock(difficulty uint) *Block {
	return NewBlock(0, genesisTimestamp, "", difficulty)
}
