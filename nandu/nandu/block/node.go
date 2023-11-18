package block

type Operator interface {
	Operate(i bool) bool
}

type Node struct {
	Op         Operator
	I          bool
	O          bool
	Next       *Node
	Identifier string
	Partner    *Node
}

func NewNode(operator Operator, identifier string) *Node {
	return &Node{
		Op:         operator,
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
	n.O = n.Op.Operate(n.I)
	if n.Next != nil {
		n.Next.Set(n.O)
	}
}
