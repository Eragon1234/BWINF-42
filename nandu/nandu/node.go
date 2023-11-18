package nandu

type Operation func(i1, i2 bool) bool

type Node struct {
	Operation  Operation
	I, O       bool
	Next       *Node
	Partner    *Node
	Identifier string
}

func NewNode(operator Operation, identifier string) *Node {
	return &Node{
		Operation:  operator,
		Identifier: identifier,
	}
}

func (n *Node) Set(i bool) {
	n.I = i
	n.Update()
	n.updatePartner()
}

func (n *Node) updatePartner() {
	if n.Partner != nil {
		n.Partner.Update()
	}
}

func (n *Node) Update() {
	n.O = n.Operation(n.I, n.Partner != nil && n.Partner.I)
	if n.Next != nil {
		n.Next.Set(n.O)
	}
}
