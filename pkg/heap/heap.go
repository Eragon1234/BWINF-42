package heap

import (
	"container/heap"
)

type Heap[T any] struct {
	h *internalHeap
}

func New[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		h: &internalHeap{
			less: func(a, b any) bool {
				return less(a.(T), b.(T))
			},
		},
	}
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.h, x)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

func (h *Heap[T]) Len() int {
	return h.h.Len()
}
