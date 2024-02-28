package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	newDim := n + m

	nums1 = append(nums1, make([]int, newDim-len(nums1))...)
	copy(nums1[len(nums1)-len(nums2):], nums2)

	mergeSort(nums1, 0, len(nums1)-1)

	fmt.Println(nums1)
}

func merge2(a []int, left, center, right int) {
	i := left
	j := center + 1
	k := 0
	b := make([]int, right-left+1)

	for i <= center && j <= right {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
			k++
		} else {
			b[k] = a[j]
			j++
			k++
		}
	}

	for i <= center {
		b[k] = a[i]
		i++
		k++
	}

	for j <= right {
		b[k] = a[j]
		j++
		k++
	}

	for k = left; k < right; k++ {
		a[k] = b[k-left]
	}
}

func mergeSort(slice []int, left, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSort(slice, left, center)
		mergeSort(slice, center+1, right)
		merge2(slice, left, center, right)
	}
}

func main() {
	nums1 := []int{1, 3, 5, 16}
	nums2 := []int{7, 14, 99, 101, 104}

	merge(nums1, len(nums1), nums2, len(nums2))
}
