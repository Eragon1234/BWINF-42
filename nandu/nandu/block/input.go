package block

type Input struct {
	next Node
}

func (b *Input) Set(i bool) {
	b.next.Set(i)
}

func (b *Input) Next() Node {
	return b.next
}

func (b *Input) SetNext(next Node) {
	b.next = next
}
