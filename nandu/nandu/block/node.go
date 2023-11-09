package block

type Node interface {
	Set(bool)
	Next() Node
	SetNext(Node)
}
