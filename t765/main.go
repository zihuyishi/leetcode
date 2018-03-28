package main

import "fmt"

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	// before is make([]int, len(row))，but slow, why?
	numPos := make(map[int]int)

	for i := 0; i < 2*n; i++ {
		numPos[row[i]] = i
	}
	swap := 0
	for i := 0; i < n; i++ {
		x, y := row[2*i], row[2*i+1]
		var move int
		var dest int
		if x > y {
			x, y = y, x
			move = 2 * i
		} else {
			move = 2*i + 1
		}

		if x%2 == 0 {
			if y == x+1 {
				continue
			}
			dest = numPos[x+1]
			numPos[y] = dest
		} else {
			dest = numPos[x-1]
			numPos[y] = dest
		}

		row[dest] = row[move]
		swap += 1
	}
	return swap
}

func main() {
	row := []int{0, 2, 3, 1, 4, 6, 5, 8, 7, 9}
	fmt.Println(minSwapsCouples(row))
}
