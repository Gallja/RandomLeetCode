package main

import (
	"fmt"
)

/*
Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

A subarray is a contiguous part of the array.
*/
func numSubarraysWithSum(nums []int, goal int) int {
	curr := 0
	ris := 0

	for i := 0; i < len(nums); i++ {
		curr += nums[i]

		if curr == goal {
			ris++
		} else if curr == goal+1 {
			break
		}
	}

	return ris
}

func subArray(nums []int, goal int) int {
	occorrenze := 0

	for len(nums) != 0 {
		occorrenze += numSubarraysWithSum(nums, goal)
		nums = nums[1:]
	}

	return occorrenze
}

func main() {
	nums := []int{0, 0, 0, 0, 0}
	fmt.Println(subArray(nums, 0))

	nums2 := []int{1, 0, 1, 0, 1}
	fmt.Println(subArray(nums2, 2))
}
