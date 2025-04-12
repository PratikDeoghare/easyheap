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
