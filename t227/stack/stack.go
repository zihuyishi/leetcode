package main

import (
	"strconv"
	"fmt"
)


func calculate(s string) int {
	buf := make([]byte, 0)
	var lastOp byte = 0
	lastNum := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			strNum := string(buf)
			num, _ := strconv.Atoi(strNum)
			buf = buf[0:0]

			if c == '*' || c == '/' {
				buf1 := make([]byte, 0)
				lastC := c
				i++
				for ; i < len(s); i++ {
					c1 := s[i]
					if c1 == ' ' {
						continue
					} else if c1 == '*' || c1 == '/' || c1 == '+' || c1 == '-' {
						strNum := string(buf1)
						num1, _ := strconv.Atoi(strNum)
						buf1 = buf1[0:0]
						if lastC == '*' {
							num = num * num1
						} else {
							num = num / num1
						}
						lastC = c1
						if c1 == '+' || c1 == '-' {
							c = c1
							break
						}
					} else {
						buf1 = append(buf1, c1)
					}
				}
				if i == len(s) && len(buf1) != 0 {
					strNum := string(buf1)
					num1, _ := strconv.Atoi(strNum)
					if lastC == '*' {
						num = num * num1
					} else {
						num = num / num1
					}
					c = '+'
				}
			}

			if c == '+' || c == '-' {
				if lastOp == '+' {
					lastNum = lastNum + num
				} else if lastOp == '-' {
					lastNum = lastNum - num
				} else if lastOp == 0 {
					lastNum = num
				}
				lastOp = c
			}
		} else {
			buf = append(buf, c)
		}
	}
	if len(buf) == 0 {
		return lastNum
	}
	strNum := string(buf)
	num, _ := strconv.Atoi(strNum)
	if lastOp == '+' {
		lastNum += num
	} else if lastOp == '-' {
		lastNum -= num
	} else if lastOp == '*' {
		lastNum *= num
	} else if lastOp == '/' {
		lastNum /= num
	} else {
		lastNum = num
	}
	return lastNum
}


func main() {
	s := "2-3*2-2/2"
	fmt.Println(calculate(s))
}