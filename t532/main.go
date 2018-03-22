package main

import (
	"fmt"
	"sort"
)

func part(nums []int, n int) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, v := range nums {
		if v < n {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

func qsort(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	n := nums[0]
	left, right := part(nums[1:], n)
	left = qsort(left)
	right = qsort(right)
	left = append(left, n)
	left = append(left, right...)
	return left
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)//sort(nums)
	numsLen := len(nums)
	ret := 0

	skip := func(arr []int, p int) int {
		v := arr[p]
		for p < len(arr)-1 {
			if arr[p+1] != v {
				break
			}
			p += 1
		}
		return p
	}
	i := 0
	for ; i < numsLen-1; {
		v := nums[i]
		for j := i+1; j < numsLen; j++ {
			if abs(nums[j] - v) > k {
				break
			}
			if abs(nums[j] - v) == k {
				ret += 1
				break
			}
		}
		i = skip(nums, i)
		i += 1
	}

	return ret
}

func main() {
	arr := []int{1, 3, 1, 5, 4}
	ret := findPairs(arr, 2)
	fmt.Println(ret)
}