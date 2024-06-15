package main

import (
	"fmt"
)

type coppia struct {
	primo   int
	secondo int
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

	for k > 0 {
		for j := 0; j < len(profits) && w >= coppie[j].primo; j++ {
			push(maxHeap, coppie[j].secondo)
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

}
