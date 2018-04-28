package main

import (
	"strings"
	"fmt"
)

func repeat(A string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += A
	}
	return ret
}

func repeatedStringMatch(A string, B string) int {
	lenA, lenB := len(A), len(B)

	if lenA >= lenB {
		if strings.Index(A, B) != -1 {
			return 1
		}
		if strings.Index(repeat(A, 2), B) != -1 {
			return 2
		}
	} else {
		s := lenB / lenA
		if lenB % lenA != 0 {
			s += 1
		}
		repA := repeat(A, s)
		if strings.Index(repA, B) != -1 {
			return s
		}
		repA += A
		if strings.Index(repA, B) != -1 {
			return s+1
		}
	}
	return -1
}

func main() {
	A := "dad"
	B := "dda"
	fmt.Print(repeatedStringMatch(A, B))
}
