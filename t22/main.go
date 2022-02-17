package main

import "fmt"

func main() {
	strs := generateParenthesis(4)
	for _, str := range strs {
		fmt.Println(str)
	}
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	results := make([]string, 0)
	for i := 0; i < n; i++ {
		inners := generateParenthesis(i)
		outers := generateParenthesis(n - 1 - i)
		for _, inner := range inners {
			for _, outer := range outers {
				str := fmt.Sprintf("(%s)%s", inner, outer)
				results = append(results, str)
			}
		}
	}
	return results
}
