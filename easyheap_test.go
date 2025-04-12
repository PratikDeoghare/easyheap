package easyheap_test

import (
	"cmp"
	"testing"

	"github.com/pratikdeoghare/easyheap"
)

type Foo struct {
	Name string
	Prio int
}

func Test_InitPushPop(t *testing.T) {
	testPrio := func(p *Foo) int {
		return p.Prio
	}
	var testSetP = func(x *Foo, p int) {
		x.Prio = p
	}

	initElems := []*Foo{
		&Foo{"Hello", 0},
		&Foo{"Hello", 10},
		&Foo{"Hello", 24},
		&Foo{"Hello", 3},
	}

	pushElems := []*Foo{
		&Foo{"&Foo1", 1},
		&Foo{"&Foo2", 2},
		&Foo{"&Foo3", 3},
		&Foo{"&Foo4", 4},
	}

	h := easyheap.NewHeap(testPrio, testSetP,
		initElems...,
	)

	for _, e := range pushElems {
		h.Push(e)
	}

	check(t, h, testPrio)

}

func Test_Remove(t *testing.T) {
	testPrio := func(p *Foo) int {
		return p.Prio
	}
	var testSetP = func(x *Foo, p int) {
		x.Prio = p
	}

	initElems := []*Foo{
		&Foo{"Hello", 0},
		&Foo{"Hello", 10},
		&Foo{"Hello", 24},
		&Foo{"Hello", 3},
	}

	pushElems := []*Foo{
		&Foo{"&Foo1", 1},
		&Foo{"&Foo2", 2},
		&Foo{"&Foo3", 3},
		&Foo{"&Foo4", 4},
	}

	h := easyheap.NewHeap(testPrio, testSetP,
		initElems...,
	)

	for _, e := range pushElems {
		h.Push(e)
	}

	t.Log(h)
	a, ok := h.Remove(initElems[0])
	t.Log(h, a, ok)
	b, ok := h.Remove(initElems[0])
	t.Log(h, b, ok)
	c, ok := h.Remove(pushElems[0])
	t.Log(h, c, ok)

	check(t, h, testPrio)
}

func Test_Fix(t *testing.T) {
	testPrio := func(p *Foo) int {
		return p.Prio
	}

	var testSetP = func(x *Foo, p int) {
		x.Prio = p
	}

	initElems := []*Foo{
		&Foo{"Hello", 0},
		&Foo{"Hello", 10},
		&Foo{"Hello", 24},
		&Foo{"Hello", 3},
	}

	pushElems := []*Foo{
		&Foo{"&Foo1", 1},
		&Foo{"&Foo2", 2},
		&Foo{"&Foo3", 3},
		&Foo{"&Foo4", 4},
	}

	h := easyheap.NewHeap(testPrio, testSetP,
		initElems...,
	)

	for _, e := range pushElems {
		h.Push(e)
	}

	t.Log(h)
	ok := h.Fix(initElems[3], -1)
	t.Log(h, ok)
	ok = h.Fix(initElems[2], -20)
	t.Log(h, ok)
	ok = h.Fix(pushElems[3], -40)
	t.Log(h, ok)

	ok = h.Fix(&Foo{}, -40)
	t.Log(h, ok)

	check(t, h, testPrio)
}

func check[P cmp.Ordered, E comparable](t *testing.T, h easyheap.Heap[P, E], testPrio func(*E) P) {
	prev := h.Pop()
	for h.Len() > 0 {
		x := h.Pop()
		t.Log(prev, x)
		if testPrio(prev) > testPrio(x) {
			t.Fatalf("expected priority of %v to be lower than %v", prev, x)
		}
		prev = x
	}
}
