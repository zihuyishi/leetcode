package main

import "fmt"

func check(s string, hasDel bool) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i += 1
			j -= 1
		} else if !hasDel {
			if check(s[i:j], true) {
				return true
			} else if check(s[i+1:j+1], true) {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	return check(s, false)
}

func main() {
	fmt.Print(validPalindrome("abcb"))
}
