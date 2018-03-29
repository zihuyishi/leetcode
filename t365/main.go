package main

import "fmt"

type Queue []int

func (q Queue) push(v int) Queue {
	return append(q, v)
}

func (q Queue) top() int {
	return q[0]
}

func (q Queue) pop() Queue {
	if len(q) > 0 {
		return q[1:]
	}
	return q
}

func (q Queue) empty() bool {
	return len(q) == 0
}

func canMeasureWater(x int, y int, z int) bool {
	if z > x + y {
		return false
	}
	if z == 0 {
		return true
	}
	reach := make([]int, x+y+1)

	if x > y {
		x, y = y, x
	}

	reach[x] = 3
	reach[y] = 2

	queue := make(Queue, 0)
	queue = queue.push(x)
	queue = queue.push(y)

	for !queue.empty() {
		node := queue.top()
		queue = queue.pop()
		if reach[node] & 1 == 1 {
			v := y - (x - node)
			if v == z {
				return true
			}
			if reach[v] == 0 {
				queue = queue.push(v)
			}
			reach[v] |= 2
			if v < x {
				reach[v] |= 1
			}
		}

		if reach[node] & 2 == 2 {
			v1 := node + x
			v2 := x - (y - node)
			if v1 == z {
				return true
			}
			if v2 == z {
				return true
			}
			if v1 < y {
				if reach[v1] == 0 {
					queue = queue.push(v1)
				}
				reach[v1] |= 2
				if v1 < x {
					reach[v1] |= 1
				}
			}
			if v2 > 0 {
				if reach[v2] == 0 {
					queue = queue.push(v2)
				}
				reach[v2] |= 3
			}
		}
	}

	return reach[z] != 0
}

func main() {
	fmt.Println(canMeasureWater(1, 1, 1))
}