package main

import (
	"sort"
	"github.com/zihuyishi/leetcode/internal"
)

func findExcept(nums []int, d int, l int, r int) (int, bool) {
	if l > r {
		return 0, false
	}
	if l == r {
		return l, nums[l] == d
	}

	p := (l + r) / 2
	if nums[p] == d {
		return p, true
	} else if nums[p] > d {
		return findExcept(nums, d, l, p-1)
	} else {
		return findExcept(nums, d, p+1, r)
	}
}


func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	nlen := len(nums)

	for i := 0; i < nlen-2; i++ {
		for j := i+1; j < nlen-1; j++ {
			cur := nums[i] + nums[j]
			need := -cur
			if need >= cur {
				if p, ok := findExcept(nums, need, j+1, nlen-1); ok {
					now := []int{nums[i],nums[j],nums[p]}
					ans = append(ans, now)
				}
			} else {
				break
			}
			for j < nlen-1 && nums[j] == nums[j+1] {
				j++
			}
		}

		for i < nlen-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return ans
}

func main() {
	arr := []int{0,0,0}

	for _, a := range threeSum(arr) {
		internal.PrintInts(a)
	}
}
