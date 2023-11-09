package block

type Block struct {
	i bool
	O bool

	operation func(i1, i2 bool) bool
	partner   *Block
	next      Node
}

func NewBlock(operation func(i1, i2 bool) bool) *Block {
	return &Block{
		operation: operation,
	}
}

func (b *Block) Set(i bool) {
	b.i = i

	b.Refresh()
}

func (b *Block) SetOperation(operation func(i1, i2 bool) bool) {
	b.operation = operation
	if b.partner != nil {
		b.partner.operation = operation
	}

	b.Refresh()
}

// SetPartner should only be called if there is no partner yet
func (b *Block) SetPartner(partner *Block) {
	if b.partner != nil {
		panic("Block already has a partner")
	}
	b.partner = partner
	partner.partner = b
}

func (b *Block) Refresh() {
	if b.partner != nil {
		b.partner.update()
	}
	b.update()
}

func (b *Block) update() {
	b.O = b.operation(b.i, b.partner != nil && b.partner.i)

	b.next.Set(b.O)
}

func (b *Block) Next() Node {
	return b.next
}

func (b *Block) SetNext(next Node) {
	b.next = next
}
