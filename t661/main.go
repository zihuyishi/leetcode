package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	h := len(M)
	A := make([][]int, h)

	for i := 0; i < h; i++ {
		w := len(M[i])
		Ai := make([]int, w)
		for j := 0; j < w; j++ {
			sum, cnt := 0, 0
			for px := -1; px < 2; px++ {
				for py := -1; py < 2; py++ {
					x := i + px
					y := j + py
					if x >= 0 && y >= 0 && x < h && y < w {
						sum += M[x][y]
						cnt += 1
					}
				}
			}
			Ai[j] = sum / cnt
		}
		A[i] = Ai
	}

	return A
}

func main() {
	fmt.Println(3 / 4)
}
