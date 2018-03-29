package main

import (
	"strconv"
	"fmt"
)

const (
	TokenNum = 0
	TokenAdd = 1
	TokenSub = 2
	TokenMul = 3
	TokenDiv = 4
)

type Token struct {
	tokenType int
	value int
}

func parseString(s string) []Token {
	tokens := make([]Token, 0)

	buf := make([]byte, 0)
	pos := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			strNum := string(buf)
			num, _ := strconv.Atoi(strNum)
			if !pos {
				num = -num
				pos = true
			}
			numToken := Token{
				tokenType: TokenNum,
				value: num,
			}
			var opType int
			if c == '+' {
				opType = TokenAdd
			} else if c == '-' {
				pos = false
				opType = TokenAdd
			} else if c == '*' {
				opType = TokenMul
			} else if c == '/' {
				opType = TokenDiv
			}
			opToken := Token{
				tokenType: opType,
			}

			tokens = append(tokens, numToken, opToken)
			buf = make([]byte, 0)
		} else {
			buf = append(buf, c)
		}
	}

	strNum := string(buf)
	num, _ := strconv.Atoi(strNum)
	if !pos {
		num = -num
	}
	numToken := Token{
		tokenType: TokenNum,
		value: num,
	}

	tokens = append(tokens, numToken)

	return tokens
}

func calcTokens(tokens []Token) int {
	t := tokens[0]
	if len(tokens) == 1 {
		return t.value
	}
	op := tokens[1]
	if op.tokenType == TokenAdd {
		return t.value + calcTokens(tokens[2:])
	} else if op.tokenType == TokenSub {
		return t.value - calcTokens(tokens[2:])
	} else {
		rv := tokens[2]
		var nv int
		if op.tokenType == TokenMul {
			nv = t.value * rv.value
		} else {
			nv = t.value / rv.value
		}
		newToken := Token {
			value: nv,
			tokenType: TokenNum,
		}
		tokens[2] = newToken
		return calcTokens(tokens[2:])
	}
}

func calculate(s string) int {
	return calcTokens(parseString(s))
}

func main() {
	s := " 1 - 2 "
	fmt.Println(calculate(s))
}