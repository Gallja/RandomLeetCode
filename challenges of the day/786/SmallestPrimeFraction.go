package main

import (
	"fmt"
)

type fraction struct {
	num int
	den int
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
					arr2 = append(arr2, fraz)
				}
			}
		}
	}

	// ordina arr2
	fmt.Println(arr2)
	arrFloat := componiArrFloat(arr2)
	fmt.Println(arrFloat)

	bubbleSort(arrFloat, arr2)
	fmt.Println(arrFloat)
	fmt.Println(arr2)

	// estrai arr2[k]
	ritorno := arr2[k-1]

	// aggiungi gli elementi di arr2[k] a arrRet
	arrRet = append(arrRet, ritorno.num)
	arrRet = append(arrRet, ritorno.den)

	return arrRet
}

func componiArrFloat(arr []fraction) []float64 {
	arrRet := []float64{}

	for i := 0; i < len(arr); i++ {
		arrRet = append(arrRet, float64(arr[i].num)/float64(arr[i].den))
	}

	return arrRet
}

func bubbleSort(array []float64, arrFraz []fraction) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
				arrFraz[i], arrFraz[j] = arrFraz[j], arrFraz[i]
			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 5}

	fmt.Println(kthSmallestPrimeFraction(arr, 3))
}
