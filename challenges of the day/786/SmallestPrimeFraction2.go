package main

import (
	"fmt"
)

type fraction struct {
	num int
	den int
	val float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	arr2 := []fraction{}
	arrRet := []int{}
	var fraz fraction

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				if i < j {
					fraz.num = arr[i]
					fraz.den = arr[j]
					fraz.val = float64(arr[i]) / float64(arr[j])
					arr2 = append(arr2, fraz)
				}
			}
		}
	}

	fmt.Println(arr2)
	// Ordinare l'array di frazioni
	bubbleSort(arr2)
	fmt.Println(arr2)

	// Comporre l'array da ritornare con il k-esimo elemento (avuto per argomento)
	arrRet = append(arrRet, arr2[k-1].num)
	arrRet = append(arrRet, arr2[k-1].den)

	return arrRet
}

func bubbleSort(array []fraction) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i].val > array[j].val {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 5}

	fmt.Println(kthSmallestPrimeFraction(arr, 3))
}
