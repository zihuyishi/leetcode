package main

import "fmt"

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minCostClimbingStairs(cost []int) int {
	high := len(cost)
	r := make([]int, high+1)
	r[0] = 0
	r[1] = 0

	for i := 2; i < high+1; i++ {
		r[i] = min(r[i-1]+cost[i-1], r[i-2]+cost[i-2])
	}
	return r[high]
}

func main() {
	stairs := []int {10, 15, 20}
	fmt.Println(minCostClimbingStairs(stairs))
}

