package main

import (
	"fmt"
)

func specialArray(nums []int) int {
	x := 0
	special := true

	for i := 0; i < len(nums); i++ {

		if nums[i] < i {
			special = false
		}

		if special {
			x = i + 1
		}
	}

	if x == 0 || x == 1 {
		return -1
	}

	return x
}

func main() {
	nums := []int{3, 5}
	fmt.Println(specialArray(nums))

	nums2 := []int{0, 0}
	fmt.Println(specialArray(nums2))

	nums3 := []int{0, 4, 3, 0, 4}
	fmt.Println(specialArray(nums3))

	nums4 := []int{3, 6, 7, 7, 0}
	fmt.Println(specialArray(nums4))
}
