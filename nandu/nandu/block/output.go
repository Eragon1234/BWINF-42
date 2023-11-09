package block

type Output struct {
	Value bool
}

func (b *Output) Set(i bool) {
	b.Value = i
}

func (b *Output) Next() Node {
	return nil
}

func (b *Output) SetNext(_ Node) {}
