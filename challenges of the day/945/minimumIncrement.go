package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	moves := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums[j]++
				moves++
			}
		}
	}

	return moves
}

func main() {
	nums := []int{1, 2, 2}

	fmt.Println(minIncrementForUnique(nums))

	nums2 := []int{3, 2, 1, 2, 1, 7}

	fmt.Println(minIncrementForUnique(nums2))
}
