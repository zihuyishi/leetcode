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

	d := make([][]bool, len1+1)

	for i := 0; i < len1+1; i++ {
		d[i] = make([]bool, len2+1)
	}

	d[0][0] = true

	for i := 0; i < len1+1; i++ {
		for j := 0; j < len2+1; j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] && d[i-1][j] {
				d[i][j] = true
			}
			if j > 0 && s2[j-1] == s3[i+j-1] && d[i][j-1] {
				d[i][j] = true
			}
		}
	}
	return d[len1][len2]
}

func main() {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbbaccc"

	fmt.Println(isInterleave(s1,s2,s3))
}