package blockchain

type Transaction struct {
	ID          string            `json:"id"`
	DataHash    string            `json:"data_hash"`
	Timestamp   int64             `json:"timestamp"`
	BlockHeight uint64            `json:"block_height"`
	Status      TransactionStatus `json:"status"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Signature   string            `json:"signature,omitempty"`
}

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusConfirmed TransactionStatus = "confirmed"
	StatusFailed    TransactionStatus = "failed"
)
