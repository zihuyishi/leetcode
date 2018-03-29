package main

import (
	"fmt"
	"github.com/zihuyishi/leetcode/internal"
)

func pow(x int, n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	} else {
		return x * pow(x, n-1)
	}
}


func getDigits(n int) []int {
	ds := make([]int, 0)

	for n > 0 {
		d := n % 10
		ds = append(ds, d)
		n = n / 10
	}
	return ds
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}

	ds := getDigits(n)

	cs := make([]int, len(ds))
	if ds[0] > 0 {
		cs[0] = 1
	}
	for i := 1; i < len(ds); i++ {
		// 当前位选择最大值，例如 123, 则拿到所有大于100的
		cs[i] = cs[i-1]
		if ds[i] > 0 {
			// 当前位选择非最大值，则可以直接计算 C(ds[i], 1) C(i, 1) * 10 ^ (i-1)
			// 例如 500,则为 (0,4) 选一，后两位选一, 其他位都有(0,9)十种选择
			// 5 * 2 * pow(10, 1)
			cs[i] += ds[i] * i * pow(10, i-1)
			if ds[i] == 1 {
				// 如果当前位值为1，则只有剩余数字的选择，如123,则有24个1
				sub := 0
				for j := i-1; j >= 0; j-- {
					sub = sub * 10 + ds[j]
				}
				cs[i] += sub + 1
			} else {
				// 如果大于1，则为当前为1时，有pow(10, len)个1
				cs[i] += pow(10, i)
			}
		}
	}

	return cs[len(ds) - 1]
}

func main() {
	t := internal.Timer{}
	t.Start()
	fmt.Println(countDigitOne(1))
	t.End()
	fmt.Printf("cost %v ms\n", t.CostInMs())
}