package main

import (
	"fmt"
	"github.com/zihuyishi/leetcode/internal"
)

type Context struct {
	x int
	y int
	d int
}

type Stack []Context

func (s Stack) push(v Context) Stack {
	return append(s, v)
}

func (s Stack) top() Context {
	return s[len(s)-1]
}

func (s Stack) pop() Stack {
	return s[0:len(s)-1]
}

func (s Stack) empty() bool {
	return len(s) == 0
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	w := len(board)
	if w == 0 {
		return false
	}
	h := len(board[0])
	if h == 0 {
		return false
	}
	reach := make([][]bool, w)
	for i := 0; i < w; i++ {
		reach[i] = make([]bool, h)
	}

	dirs := [][]int {
		{-1,0},
		{0,-1},
		{1,0},
		{0,1},
	}

	stack := make(Stack, 0)
	dep := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if board[i][j] == word[0] {
				reach[i][j] = true
				stack = stack.push(Context{x: i, y: j, d: 0})
				dep = 1

				for !stack.empty() {
					cur := stack.top()
					stack = stack.pop()
					if dep == len(word) {
						return true
					}
					if cur.d >= 4 {
						reach[cur.x][cur.y] = false
						dep--
						continue
					}
					d := dirs[cur.d]

					cur.d += 1
					stack = stack.push(cur)

					x, y := cur.x + d[0], cur.y + d[1]
					if x >= 0 && y >= 0 && x < w && y < h &&
						!reach[x][y] && word[dep] == board[x][y] {
							dep++
							if dep == len(word) {
								return true
							}
							reach[x][y] = true
							stack = stack.push(Context{x: x, y: y, d: 0})
					}
				}

				dep = 0
				reach[i][j] = false
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A'},
		{'B'},
		{'C'},
	}
	s := ""
	t := internal.Timer{}
	t.Start()
	fmt.Println(exist(board, s))
	t.End()
	fmt.Printf("cost %v ms", t.CostInMs())
}