package heap

// internalHeap is a wrapper around a slice implementing heap.Interface.
type internalHeap struct {
	s    []any
	less func(any, any) bool
}

func (h internalHeap) Len() int {
	return len(h.s)
}

func (h internalHeap) Less(i, j int) bool {
	return h.less(h.s[i], h.s[j])
}

func (h internalHeap) Swap(i, j int) {
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func (h *internalHeap) Push(x any) {
	h.s = append(h.s, x)
}

func (h *internalHeap) Pop() any {
	old := h.s
	n := len(old)
	x := old[n-1]
	h.s = old[:n-1]
	return x
}
