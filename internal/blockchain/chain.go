package blockchain

type Chain struct {
	Blocks     []*Block
	difficulty uint
}

func NewChain(difficulty uint) (*Chain, error) {
	genesisBlock := NewGenesisBlock(difficulty)
	return &Chain{
		Blocks:     []*Block{genesisBlock},
		difficulty: difficulty,
	}, nil

}
