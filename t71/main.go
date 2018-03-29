package main

import (
	"strings"
	"fmt"
)

func simplifyPath(path string) string {
	stack := make([]string,0)

	paths := strings.Split(path, "/")

	for _, p := range paths {
		if p == "." || p == "" {
			continue
		} else if p == ".." {
			if len(stack) > 0 {
				stack = stack[0:len(stack)-1]
			}
		} else {
			stack = append(stack, p)
		}
	}

	ret := "/"
	ret += strings.Join(stack, "/")
	return ret
}

func main() {
	path := "/.."
	fmt.Println(simplifyPath(path))
}