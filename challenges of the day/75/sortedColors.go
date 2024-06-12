package main

import (
	"fmt"
)

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	half := len(nums) / 2
	left := make([]int, half)
	right := make([]int, len(nums)-half)

	copy(left, (nums)[:half])
	copy(right, (nums)[half:])

	sortColors(left)
	sortColors(right)

	merged := merge(left, right)
	for i := 0; i < len(nums); i++ {
		(nums)[i] = merged[i]
	}
}

func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)
}
