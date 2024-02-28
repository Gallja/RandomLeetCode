package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	ris := make([]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ris = append(ris, i)
				ris = append(ris, j)
			}
		}
	}

	return ris
}

func main() {
	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}
