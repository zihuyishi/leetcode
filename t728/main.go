package main

import (
	"github.com/zihuyishi/leetcode/internal"
	"fmt"
)


func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	for i := left; i < right+1; i++ {
		pass := true
		for n := i; n != 0; n /= 10 {
			if (n % 10) == 0 || i % (n % 10) != 0 {
				pass = false
				break
			}
		}
		if pass {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	t := &internal.Timer{

	}
	t.Start()
	ret := selfDividingNumbers(1, 10000)
	t.End()
	fmt.Println(t.CostInMs())
	internal.PrintInts(ret)
}