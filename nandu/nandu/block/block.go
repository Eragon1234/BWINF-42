package block

type Block struct {
	i bool
	O bool

	operation func(i1, i2 bool) bool
	SetOutput func(o bool)
	partner   *Block
}

func NewNoBlock() *Block {
	return &Block{
		operation: NoBlock,
		SetOutput: func(o bool) {},
	}
}

func (b *Block) I() bool {
	return b.i
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
	if b.partner == nil {
		b.O = b.operation(b.i, false)
	} else {
		b.O = b.operation(b.i, b.partner.I())
	}

	b.SetOutput(b.O)
}
