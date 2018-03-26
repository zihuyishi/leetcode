package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)

	if len1 + len2 != len3 {
		return false
	}

	if len3 == 0 {
		return true
	}

	if len1 > 0 && s1[0] == s3[0] {
		if isInterleave(s1[1:], s2, s3[1:]) {
			return true
		}
	}

	if len2 > 0 && s2[0] == s3[0] {
		if isInterleave(s1, s2[1:], s3[1:]) {
			return true
		}
	}

	return false
}

func main() {
	s1, s2, s3 := "a", "b", "ab"

	fmt.Println(isInterleave(s1,s2,s3))
}