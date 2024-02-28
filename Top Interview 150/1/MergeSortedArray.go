package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	newDim := n + m

	nums1 = append(nums1, make([]int, newDim-len(nums1))...)
	copy(nums1[len(nums1)-len(nums2):], nums2)
}

func main() {
	nums1 := []int{1, 3, 5, 16}
	nums2 := []int{7, 14, 99, 101, 104}

	merge(nums1, len(nums1), nums2, len(nums2))

	fmt.Println(nums1)
}
