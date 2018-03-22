package internal

import "fmt"

func PrintInts(a []int) {
	for _, v := range a {
		fmt.Printf("%v,", v)
	}
	fmt.Println()
}
