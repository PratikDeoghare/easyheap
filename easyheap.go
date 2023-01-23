package easyheap

import "container/heap"

type pheap[E any] struct {
	elems []E
	less  func(E, E) bool
}

func (h pheap[E]) Len() int           { return len(h.elems) }
func (h pheap[E]) Less(i, j int) bool { return h.less(h.elems[i], h.elems[j]) }
func (h pheap[E]) Swap(i, j int)      { h.elems[i], h.elems[j] = h.elems[j], h.elems[i] }

func (h *pheap[E]) Push(x any) {
	h.elems = append(h.elems, x.(E))
}

func (h *pheap[E]) Pop() any {
	old := h.elems
	n := len(old)
	x := old[n-1]
	h.elems = old[0 : n-1]
	return x
}

type Heap[E any] struct {
	h *pheap[E]
}

func NewHeap[E any](less func(E, E) bool, es ...E) Heap[E] {
	h := pheap[E]{
		less: less,
	}

	for _, e := range es {
		h.elems = append(h.elems, e)
	}

	heap.Init(&h)

	return Heap[E]{
		h: &h,
	}
}

func (e *Heap[E]) Push(x E) {
	heap.Push(e.h, x)
}

func (e *Heap[E]) Pop() E {
	return heap.Pop(e.h).(E)
}

func (e Heap[E]) Len() int {
	return e.h.Len()
}

//
//func (e Heap[E]) Remove(i int) E {
//	return heap.Remove(e.h, i).(E)
//}
//
//func (e Heap[E]) Fix(i int) {
//	heap.Fix(e.h, i)
//}
