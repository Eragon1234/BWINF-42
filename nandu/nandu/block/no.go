package block

type NoBlock struct{}

func (n *NoBlock) Operate(_ bool) bool {
	return false
}
