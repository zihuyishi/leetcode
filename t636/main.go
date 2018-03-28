package main

import (
	"github.com/zihuyishi/leetcode/internal"
	"strings"
	"strconv"
)

type Log struct {
	id   int
	time int
	t    int
}
type Stack []*Log

func (s Stack) push(v *Log) Stack {
	return append(s, v)
}

func (s Stack) top() *Log {
	if len(s) != 0 {
		return s[len(s)-1]
	}
	return nil
}

func (s Stack) pop() Stack {
	return s[:len(s)-1]
}

func parseLog(s string) *Log {
	values := strings.Split(s, ":")
	id, err := strconv.Atoi(values[0])
	if err != nil {
		return nil
	}
	var t int
	if values[1] == "start" {
		t = 0
	} else {
		t = 1
	}
	time, err := strconv.Atoi(values[2])
	if err != nil {
		return nil
	}
	return &Log{
		id:   id,
		t:    t,
		time: time,
	}
}

func exclusiveTime(n int, logs []string) []int {
	spent := make([]int, n)
	stack := make(Stack, n)

	var lastLog *Log

	for _, sLog := range logs {
		log := parseLog(sLog)

		curLog := stack.top()
		fix := 0
		if log.t == 1 {
			stack = stack.pop()
			if lastLog.t == 0 {
				fix = 1
			}
			spent[curLog.id] += log.time - lastLog.time + fix
		} else {
			if lastLog != nil && curLog != nil {
				if lastLog.t == 1 {
					fix = -1
				}
				spent[curLog.id] += log.time - lastLog.time + fix
			}
			stack = stack.push(log)
		}
		lastLog = log
	}

	return spent
}

func main() {
	logs := []string{
		"0:start:0",
		"0:end:0",
		"1:start:1",
		"1:end:1",
		"2:start:2",
		"2:end:2",
		"2:start:10",
		"2:end:10",
	}
	internal.PrintInts(exclusiveTime(3, logs))
}
