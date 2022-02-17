package main

import "fmt"

func main() {
	result := divide(-3, -3)
	fmt.Printf("%d", result)
}

// TODO: handle overflow
func divide(dividend int, divisor int) int {
	neg := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		neg = true
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if dividend < divisor {
		return 0
	}
	result := 1
	cur := divisor
	for {
		if cur+cur <= dividend {
			result += result
			cur += cur
		} else if cur+divisor <= dividend {
			result += 1
			cur += divisor
		} else {
			break
		}
	}

	if neg {
		return -result
	} else {
		return result
	}
}
