package main

import "fmt"

func longestPalindrome(s string) int {
	freqs := make([]bool, 256)

	ret := 0
	for _, c := range s {
		if freqs[c] {
			ret += 2
			freqs[c] = false
		} else {
			freqs[c] = true
		}
	}

	if ret != len(s) {
		ret += 1
	}

	return ret
}

func main() {
	s := "aaabb"
	fmt.Println(longestPalindrome(s))
}
