package main

import "fmt"

func canTransform(start string, end string) bool {
	s_buf := ([]uint8)(start)
	e_buf := ([]uint8)(end)

	slen := len(s_buf)

	for i := 0; i < slen; i++ {
		if s_buf[i] == e_buf[i] {
			continue
		}

		if i == slen - 1 {
			return false
		}

		if s_buf[i] == 'X' && e_buf[i] == 'L' {
			for j := i+1; j < slen; j++ {
				if s_buf[j] == 'R' {
					return false
				} else if s_buf[j] == 'L' {
					s_buf[i], s_buf[j] = s_buf[j], s_buf[i]
					break
				} else if j == slen-1 {
					return false
				}
			}
		} else if s_buf[i] == 'R' && e_buf[i] == 'X' {
			for j := i+1; j < slen; j++ {
				if s_buf[j] == 'L' {
					return false
				} else if s_buf[j] == 'X' {
					s_buf[i], s_buf[j] = s_buf[j], s_buf[i]
					break
				} else if j == slen - 1 {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "RXL"
	e := "XRL"
	fmt.Println(canTransform(s, e))
}