package heap

import (
	"container/heap"
)

type Heap[T any] struct {
	h *internalHeap
}

// New returns a new heap with the given less function.
//
// If the less function returns true,
// the first element will be popped before the second.
func New[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		h: &internalHeap{
			less: func(a, b any) bool {
				return less(a.(T), b.(T))
			},
		},
	}
}

// Push adds x to the heap.
func (h *Heap[T]) Push(x T) {
	heap.Push(h.h, x)
}

// Pop removes and returns the minimum element (according to less).
func (h *Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

// Len returns the number of elements in the heap.
func (h *Heap[T]) Len() int {
	return h.h.Len()
}
