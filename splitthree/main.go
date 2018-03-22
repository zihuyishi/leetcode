package main

import "fmt"

func solution(arr []int) int {
	alen := len(arr)

	mins := make([]int, alen)
	d := make([]int, alen)

	for i:= 1; i < alen-1; i++ {
		v := arr[i]
		if mins[i-1] > v || mins[i-1] == 0 {
			mins[i] = v
		} else {
			mins[i] = mins[i-1]
		}

		if i > 2 {
			d[i] = v + mins[i-2]
		}
	}

	ret := d[3]
	for i := 3; i < alen-1; i++ {
		if d[i] < ret {
			ret = d[i]
		}
	}
	return ret
}

func main() {
	a := []int{5,3,2,1,2,2,3,4,5}
	fmt.Println(solution(a))
}
