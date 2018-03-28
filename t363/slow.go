package main

import "fmt"

func maxSumSubmatrix(matrix [][]int, k int) int {
	result := matrix[0][0]

	w := len(matrix)
	h := len(matrix[0])
	sums := make([][]int, w)
	for i := 0; i < w; i++ {
		line := make([]int, h)
		acc := 0
		for j := 0; j < h; j++ {
			acc += matrix[i][j]
			if i >= 1 {
				line[j] = sums[i-1][j] + acc
			} else {
				line[j] = acc
			}
		}
		sums[i] = line
	}

	for s := 1; s <= w; s++ {
		for t := 1; t <= h; t++ {
			for i := s-1; i < w; i++ {
				for j := t-1; j < h; j++ {
					now := sums[i][j]
					if i >= s {
						now -= sums[i-s][j]
					}
					if j >= t {
						now -= sums[i][j-t]
					}
					if i >= s && j >= t {
						now += sums[i-s][j-t]
					}
					if (result > k || now > result) && now <= k {
						result = now
					}
				}
			}
		}
	}

	return result
}

func main() {
	k := 2
	matrix := [][]int {
		{-4,-2,-1},
		{1,7,1},
	}

	fmt.Println(maxSumSubmatrix(matrix, k))
}