package block

type NoOp struct{}

func (o *NoOp) Operate(i bool) bool {
	return i
}
