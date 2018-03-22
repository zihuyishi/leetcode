package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	numLen := len(nums)
	p := 1
	i := 1
	for i < numLen {
		if nums[i] != nums[p-1] {
			nums[p] = nums[i]
			p += 1
		}
		i += 1
	}
	return p
}

func main() {
	arr := []int{1,1,2,3,3,4,5,5,6}
	fmt.Println(removeDuplicates(arr))
}