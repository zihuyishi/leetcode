package main

import "github.com/zihuyishi/leetcode/internal"


func countSmaller(nums []int) []int {
	nlen := len(nums)
	ret := make([]int, len(nums))
	if nlen == 0 {
		return ret
	}
	bigs := make([]int, len(nums))

	ret[nlen-1] = 0
	bigs[nlen-1] = nlen-1

	for i := len(nums)-2; i >= 0; i-- {
		v := nums[i]
		if v > nums[bigs[i+1]] {
			bigs[i] = i
			ret[i] = nlen - i - 1
		} else {
			bigs[i] = bigs[i+1]
			c := 0
			for j := i+1; j < nlen; j++ {
				if v > nums[bigs[j]] {
					c += nlen - j
					break
				}
				if v > nums[j] {
					c += 1
				}
			}
			ret[i] = c
		}
	}

	return ret
}

func main() {
	nums := []int{94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41}
	ret := countSmaller(nums)
	internal.PrintInts(ret)
}