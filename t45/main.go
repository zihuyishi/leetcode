package main

import "fmt"

func jump(nums []int) int {
	ret := 0

	p := 0
	numLen := len(nums)
	for p < numLen {
		if p == numLen - 1 {
			break
		}
		step := nums[p]
		max := 0
		next := p+1
		for i := p+1; i <= p + step; i++ {
			if i >= numLen-1 {
				next = numLen - 1
				break
			}
			if nums[i] + i >= max {
				max = nums[i] + i
				next = i
			}
		}
		// jump to max
		p = next
		ret += 1
		// fmt.Println(p)
	}

	return ret
}

func main() {
	arr := []int{}
	ret := jump(arr)
	fmt.Printf("result is %v\n", ret)
}