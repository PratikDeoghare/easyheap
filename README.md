# Easyheaps 

Easyheaps asks you to provide only the comparison function for the items you
want to put in the heap. It exposes only Push, Pop and Len. This is sufficient for
everyday coding. `container/heap` packge allows you to do much more so you can
switch to that if needed. 


``` 

type foo struct {
	name string
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
``` 


This package exists because I don't fancy typing this again. 

``` 
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

``` 


