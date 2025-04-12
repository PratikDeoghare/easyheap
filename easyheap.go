package easyheap

import (
	"cmp"
	"container/heap"
	"fmt"
)

type hp[P cmp.Ordered, E any] struct {
	elems []*E
	getp  func(e *E) P
	setp  func(e *E, p P)
	idx   map[*E]int
}

func (h hp[P, E]) Len() int           { return len(h.elems) }
func (h hp[P, E]) Less(i, j int) bool { return h.getp(h.elems[i]) < h.getp(h.elems[j]) }
func (h hp[P, E]) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
	h.idx[h.elems[i]] = i
	h.idx[h.elems[j]] = j
}

func (h *hp[P, E]) Push(x any) {
	y := x.(*E)
	h.idx[y] = len(h.elems)
	h.elems = append(h.elems, y)
}

func (h *hp[P, E]) Pop() any {
	old := h.elems
	n := len(old)
	x := old[n-1]
	delete(h.idx, x)
	h.elems = old[0 : n-1]
	return x
}

type Heap[P cmp.Ordered, E comparable] struct {
	h *hp[P, E]
}

func NewHeap[P cmp.Ordered, E comparable](getp func(e *E) P, setp func(e *E, p P), es ...*E) Heap[P, E] {
	h := hp[P, E]{
		idx:  make(map[*E]int),
		getp: getp,
		setp: setp,
	}

	for i, e := range es {
		h.idx[e] = i
		h.elems = append(h.elems, e)
	}

	heap.Init(&h)

	return Heap[P, E]{
		h: &h,
	}
}

func (e *Heap[P, E]) Push(x *E) {
	heap.Push(e.h, x)
}

func (e *Heap[P, E]) Pop() *E {
	return heap.Pop(e.h).(*E)
}

func (e Heap[P, E]) Len() int {
	return e.h.Len()
}

func (e Heap[P, E]) Remove(x *E) (*E, bool) {
	a, ok := e.h.idx[x]
	if !ok {
		return nil, false
	}
	return heap.Remove(e.h, a).(*E), true
}

func (e Heap[P, E]) Fix(x *E, priority P) bool {
	e.h.setp(x, priority)
	a, ok := e.h.idx[x]
	if !ok {
		return false
	}
	heap.Fix(e.h, a)
	return true
}

func (e Heap[P, E]) String() string {
	/*
		for e, idx := range e.h.idx {
			fmt.Println(e)
			for x := range idx {
				fmt.Println("\t", x)
			}
		}
	*/

	return fmt.Sprintf("%d %d", len(e.h.elems), len(e.h.idx))
}
