package block

type Operator interface {
	Operate(i bool) bool
}

type Node struct {
	Op           Operator
	I            bool
	O            bool
	Next         *Node
	Identifier   string
	Dependencies []*Node
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
	n.updateDependencies()
}

func (n *Node) updateDependencies() {
	for _, dep := range n.Dependencies {
		dep.Update()
	}
}

func (n *Node) Update() {
	n.O = n.Op.Operate(n.I)
	if n.Next != nil {
		n.Next.Set(n.O)
	}
}
