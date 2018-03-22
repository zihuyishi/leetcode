package main

import "fmt"

type bound struct {
	left int
	right int
}

func findShortestSubArray(nums []int) int {
	freqs := make(map[int]int)

	numsLen := len(nums)
	bounds := make([]*bound, 50000)

	for i, v := range nums {
		freqs[v] += 1
		now := nums[i]

		var b = bounds[now]

		if b == nil {
			b = &bound {
				left: numsLen,
				right: -1,
			}
			bounds[now] = b
		}

		if b.left > i {
			b.left = i
		}

		if b.right < i {
			b.right = i
		}
	}

	var maxs []int
	topper := 0

	for k, v := range freqs {
		if v > topper {
			maxs = make([]int, 0)
			topper = v
		}

		if v == topper {
			maxs = append(maxs, k)
		}
	}

	ret := numsLen

	for _, v := range maxs {
		bound := bounds[v]
		dist := bound.right - bound.left + 1
		if dist < ret {
			ret = dist
		}
	}

	return ret
}

func main() {
	arr := []int{}
	ret := findShortestSubArray(arr)

	fmt.Printf("result is %v", ret)
}

func printMap(m *map[int]int) {
	for k, v := range *m {
		fmt.Printf("%v has %v times\n", k, v)
	}
}

func printBounds(b *map[int]*bound) {
	for k, v := range *b {
		fmt.Printf("%v range is (%v, %v)\n", k, v.left, v.right)
	}
}