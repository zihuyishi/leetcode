package main

import "fmt"

type record struct {
	v int
	x1 int
	x2 int
	x3 int
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	nlen := len(nums)
	sums := make([]int, len(nums))
	d1 := make([]*record, len(nums))
	d2 := make([]*record, len(nums))
	d3 := make([]*record, len(nums))

	for i := 0; i < k; i++ {
		sums[0] += nums[i]
	}
	d1[0] = &record {
		v: sums[0],
		x1: 0,
	}

	for i := 1; i < nlen-k+1; i++ {
		sums[i] = sums[i-1] - nums[i-1] + nums[i+k-1]
		if d1[i-1] != nil && d1[i-1].v >= sums[i] {
			d1[i] = d1[i-1]
		} else {
			d1[i] = &record {
				v: sums[i],
				x1: i,
			}
		}
		if i >= k {
			if d2[i-1] != nil && d2[i-1].v >= d1[i-k].v + sums[i] {
				d2[i] = d2[i-1]
			} else {
				d2[i] = &record{
					v: d1[i-k].v + sums[i],
					x1: d1[i-k].x1,
					x2: i,
				}
			}
		}
		if i >= 2 * k {
			if d3[i-1] != nil && d3[i-1].v >= d2[i-k].v + sums[i] {
				d3[i] = d3[i-1]
			} else {
				d3[i] = &record{
					v: d2[i-k].v + sums[i],
					x1: d2[i-k].x1,
					x2: d2[i-k].x2,
					x3: i,
				}
			}
		}
	}

	return []int{d3[nlen-k].x1, d3[nlen-k].x2, d3[nlen-k].x3}
}

func main() {
	arr := []int{1,2,3,2,6,7,5,1}
	k := 2
	fmt.Println(maxSumOfThreeSubarrays(arr, k))
}