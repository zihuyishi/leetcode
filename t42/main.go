package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	lhighest := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		if i > 0 && height[lhighest[i-1]] > height[i] {
			lhighest[i] = lhighest[i-1]
		} else {
			lhighest[i] = i
		}
	}

	rhighest := make([]int, len(height))

	for i := len(height)-1; i >= 0; i-- {
		if i < len(height)-1 && height[rhighest[i+1]] > height[i] {
			rhighest[i] = rhighest[i+1]
		} else {
			rhighest[i] = i
		}
	}

	ret := 0
	mid := lhighest[len(height)-1]
	for i := mid-1; i >= 0; i-- {
		ret += height[lhighest[i]] - height[i]
	}
	for i := mid+1; i < len(height); i++ {
		ret += height[rhighest[i]] - height[i]
	}
	return ret
}

func main() {
	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(arr))
}