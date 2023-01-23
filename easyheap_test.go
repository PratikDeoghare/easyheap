package easyheap_test

import (
	"fmt"
	"testing"

	"github.com/pratikdeoghare/easyheap"
)

type foo struct {
	name string
}

func TestName(t *testing.T) {

	type lexfoo foo

	h := easyheap.NewHeap[*lexfoo](
		func(a, b *lexfoo) bool {
			return a.name < b.name
		},
		&lexfoo{"today"},
		&lexfoo{"is"},
		&lexfoo{"the"},
		&lexfoo{"day"},
	)

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	h.Push(&lexfoo{"di"})

	hh := easyheap.NewHeap[*foo](
		func(a, b *foo) bool { return len(a.name) > len(b.name) },
		&foo{"today"},
		&foo{"is"},
		&foo{"the"},
		&foo{"day"},
	)

	for hh.Len() > 0 {
		fmt.Println(hh.Pop())
	}
}

func ExampleReverse() {

	type lexfoo foo

	h := easyheap.NewHeap[*lexfoo](
		func(a, b *lexfoo) bool {
			return a.name < b.name
		},
		&lexfoo{"today"},
		&lexfoo{"is"},
		&lexfoo{"the"},
		&lexfoo{"day"},
	)

	for h.Len() > 0 {
		fmt.Println(h.Pop().name)
	}
	// Output:
	//day
	//is
	//the
	//today
}
