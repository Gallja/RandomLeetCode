package main

import (
	"fmt"
)

func specialArray(nums []int) int {
	x := 0
	special := true

	for i := 0; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			if nums[i] < i {
				special = false
				break
			}
		}

		if special {
			x = i + 1
		}
	}

	if x == 0 {
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
}
