package block

type Block struct {
	Operation func(i1, i2 bool) bool
	Partner   *Node
}

func NewBlock(operation func(i1, i2 bool) bool) *Block {
	return &Block{
		Operation: operation,
	}
}

func (b Block) Operate(i bool) bool {
	return b.Operation(i, b.Partner.I)
}
