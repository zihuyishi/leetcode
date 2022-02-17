package main

import (
	"fmt"
	"strings"
)

func main() {
	a := multiply("1", "12")
	fmt.Println(a)
}

func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	result := make([]uint8, len1+len2+1)
	for i := 0; i < len1; i++ {
		var added uint8 = 0
		for j := 0; j < len2; j++ {
			v1 := num1[len1-i-1] - '0'
			v2 := num2[len2-j-1] - '0'
			v := v1*v2 + added + result[i+j]
			remainder := v % 10
			result[i+j] = remainder
			added = v / 10
		}
		if added != 0 {
			result[i+len2] += added
		}
	}
	sb := strings.Builder{}
	start := false
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == 0 && !start {
			continue
		}
		start = true
		sb.WriteByte(result[i] + '0')
	}
	if sb.Len() == 0 {
		sb.WriteByte('0')
	}
	return sb.String()
}
