package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := (len(nums1) + len(nums2) + 1) / 2
	m1 := len(nums1) / 2
	m2 := len(nums2) / 2

	return 0.0
}

func main() {
	s1 := []int{1, 3, 6, 7, 8}
	s2 := []int{5, 8, 9, 10, 12, 13, 14, 15, 16}
	// 1,2,5,6,7,8,8,9,10,12,13,14,15,16
	fmt.Print(findMedianSortedArrays(s1, s2))
}
