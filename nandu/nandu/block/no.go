package block

type NoBlock struct {
}

func (n *NoBlock) Set(_ bool) {}

func (n *NoBlock) Next() Node {
	return nil
}

func (n *NoBlock) SetNext(_ Node) {}
