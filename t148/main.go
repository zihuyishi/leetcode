package main

import (
	"fmt"
	"github.com/zihuyishi/leetcode/internal"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pivot(list *ListNode, v int) (*ListNode, *ListNode) {
	if list == nil {
		return nil, nil
	}

	var left, right *ListNode

	curLeft := left
	curRight := right

	for list != nil {
		next := list.Next
		list.Next = nil
		if list.Val < v {
			if left == nil {
				left = list
				curLeft = list
			} else {
				curLeft.Next = list
				curLeft = list
			}
		} else {
			if right == nil {
				right = list
				curRight = list
			} else {
				curRight.Next = list
				curRight = list
			}
		}
		list = next
	}

	return left, right
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func qsortList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	mid := head

	var left, right *ListNode

	curLeft := left
	curRight := right

	head = head.Next
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < mid.Val {
			if left == nil {
				left = head
				curLeft = head
			} else {
				curLeft.Next = head
				curLeft = head
			}
		} else {
			if right == nil {
				right = head
				curRight = head
			} else {
				curRight.Next = head
				curRight = head
			}
		}
		head = next
	}

	left, ltail := qsortList(left)
	right, rtail := qsortList(right)

	if left == nil {
		left = mid
	} else {
		ltail.Next = mid
	}
	if right == nil {
		rtail = mid
	}
	mid.Next = right

	return left, rtail
}

func sortList(head *ListNode) *ListNode {
	ret, _ := qsortList(head)
	return ret
}

func (list *ListNode) len() int {
	cur := list
	ret := 0
	for cur != nil {
		ret += 1
		cur = cur.Next
	}
	return ret
}



func makeList(arr []int) *ListNode {
	alen := len(arr)
	if alen == 0 {
		return nil
	}

	head := &ListNode{
		Val: arr[0],
		Next: nil,
	}

	cur := head

	for i := 1; i < alen; i++ {
		node := &ListNode{
			Val: arr[i],
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	return head
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%v,", list.Val)
		list = list.Next
	}
	fmt.Println()
}


func main() {
	t := &internal.Timer{}
	s := 1000000
	arr := make([]int, s)
	for i := 0; i < s; i++ {
		arr[i] = rand.Int()
	}
	l := makeList(arr)
	t.Start()
	l = sortList(l)
	t.End()
	fmt.Printf("cost %v ms\n", t.CostInMs())
}