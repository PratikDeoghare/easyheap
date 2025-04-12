# Easyheap
[![Go](https://github.com/PratikDeoghare/easyheap/actions/workflows/go.yml/badge.svg)](https://github.com/PratikDeoghare/easyheap/actions/workflows/go.yml)

```
package main

import (
	"fmt"

	"github.com/pratikdeoghare/easyheap"
)

func main() {

	type Player struct {
		Name string
	}

	scores := make(map[string]int)

	getPriority := func(x *Player) int {
		return scores[x.Name]
	}

	setPriority := func(x *Player, score int) {
		scores[x.Name] = score
	}

	players := []*Player{
		&Player{"Jan"},
		&Player{"Johnny"},
		&Player{"Janardan"},
		&Player{"Tara"},
		&Player{"Ram"},
		&Player{"Pam"},
	}

	for i, p := range players {
		setPriority(p, -i*200)
	}

	h := easyheap.NewHeap(getPriority, setPriority, players...)

	for h.Len() > 0 {
		x := h.Pop()
		fmt.Println(x)
	}

}
```




This package exists because I don't fancy typing this all the time. 

``` 
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

``` 


