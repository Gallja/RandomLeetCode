package main

import (
	"fmt"
)

type coppia struct {
	capitale int
	profitto int
}

type heap []int

func lenHeap(maxHeap heap) int {
	return len(maxHeap)
}

func push(maxHeap *heap, x int) {
	*maxHeap = append(*maxHeap, x)
}

func pop(maxheap *heap) int {
	bkcp := *maxheap
	len := len(*maxheap)

	last := bkcp[len-1]
	*maxheap = bkcp[0 : len-1]

	return last
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	coppie := make([]coppia, len(profits))
	maxHeap := &heap{}

	for i := 0; i < len(profits); i++ {
		coppie[i] = coppia{capital[i], profits[i]}
	}

	/*
		sort.Slice(coppie, func(i, j int) bool {
			return coppie[i].capitale < coppie[j].capitale
		})
	*/

	j := 0

	for k > 0 {
		for j < len(profits) && w >= coppie[j].capitale {
			push(maxHeap, coppie[j].profitto)

			j++
		}

		if lenHeap(*maxHeap) == 0 {
			break
		}

		w += pop(maxHeap)
		k--
	}

	return w
}

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))

	k2 := 3
	w2 := 0
	profits2 := []int{1, 2, 3}
	capital2 := []int{0, 1, 2}
	fmt.Println(findMaximizedCapital(k2, w2, profits2, capital2))

	fmt.Println(findMaximizedCapital(10, w, profits, capital2))
}
