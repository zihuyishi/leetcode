package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxs := make([]int, len(nums))
	cur := 1
	maxs[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			maxs[i] = maxs[i-1] + 1
		} else {
			maxs[i] = 1
		}
		if maxs[i] > cur {
			cur = maxs[i]
		}
	}
	return cur
}

func main() {
	arr := []int{1,2,3,2,4,6,8,9,2}
	fmt.Print(findLengthOfLCIS(arr))
}
