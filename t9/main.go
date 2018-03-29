package main

import (
	"math"
	"fmt"
)


func getn(x int, n int) int {
	x = x / int(math.Pow10(n))
	x = x % 10
	return x
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xlen := int(math.Floor(math.Log10(float64(x)))) + 1
	for i := 0; i < xlen/2; i++ {
		if getn(x, i) != getn(x, xlen-i-1) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(11211))
}