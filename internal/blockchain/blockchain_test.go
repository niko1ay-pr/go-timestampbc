package blockchain

import "testing"

func TestNewChain(t *testing.T) {
	genesisDifficulty := uint(3)
	chain, err := NewChain(genesisDifficulty)
	if err != nil {
		t.Fatalf("NewChain() error = %v", err)
	}
	if len(chain.Blocks) != 1 {
		t.Fatalf("NewChain() blocks length = %d, want 1", len(chain.Blocks))
	}
	genesis := chain.Blocks[0]
	if genesis.Index != 0 {
		t.Errorf("genesis.Index = %d, want 0", genesis.Index)
	}
	if genesis.Difficulty != genesisDifficulty {
		t.Errorf("genesis.Difficulty = %d, want %d", genesis.Difficulty, genesisDifficulty)
	}
	if genesis.PrevHash != "" {
		t.Errorf("genesis.PrevHash = %q, want empty string", genesis.PrevHash)
	}
	if genesis.Timestamp != genesisTimestamp {
		t.Errorf("genesis.Timestamp = %d, want %d", genesis.Timestamp, genesisTimestamp)
	}
	if genesis.Transactions != nil {
		t.Errorf("genesis.Transactions = %#v, want nil", genesis.Transactions)
	}
}
