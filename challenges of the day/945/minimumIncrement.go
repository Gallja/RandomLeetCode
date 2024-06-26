package main

import (
	"fmt"
	"sort"
)

// This function's right, but time complexity's O(n^2)
// This solution is not acceptable
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

// This is the new version of minIncrementForUnique function
// Time complexity's less than O(n^2) and it is an acceptable solution
func minIncrementForUnique2(nums []int) int {
	moves := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			for nums[i+1] <= nums[i] {
				nums[i+1]++
				moves++
				if nums[i] == nums[i+1] {
					nums[i+1]++
					moves++
				}
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

	nums3 := []int{3, 2, 1, 2, 1, 7}
	fmt.Println(minIncrementForUnique2(nums3))

	nums4 := []int{1, 2, 2}
	fmt.Println(minIncrementForUnique2(nums4))

	nums5 := []int{2, 2, 2, 2, 0}
	fmt.Println(minIncrementForUnique2(nums5))
}
