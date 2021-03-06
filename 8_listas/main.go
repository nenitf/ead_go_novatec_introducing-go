package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// lista com apontamento para o próximo item e o anterior
	// 1 -> 2 -> 3
	// 1 <- 2 <- 3

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	for e := x.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}
