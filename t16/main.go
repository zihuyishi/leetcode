package main

import (
	"sort"
	"fmt"
)

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func threeSumClosest(nums []int, target int) int {
	nlen := len(nums)

	sort.Ints(nums)
	ret := nums[0]+nums[1]+nums[2]
	for i := 0; i < nlen-2; i++ {
		j := i+1
		k := nlen-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum - target) < abs(ret - target) {
				ret = sum
			}
			if sum == target {
				return ret
			} else if sum < target {
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return ret
}

func main() {
	arr := []int{-1,2,1,-4}
	fmt.Println(threeSumClosest(arr, -100))
}
