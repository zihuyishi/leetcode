package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func findNext(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	cur := node.Next
	for cur != nil {
		if cur.Val == node.Val {
			cur = cur.Next
			if cur == nil {
				return nil
			}
		} else if node.Next == cur {
			return node
		} else {
			node = cur
			cur = node.Next
		}
	}
	return node
}

func deleteDuplicates(head *ListNode) *ListNode {
	head = findNext(head)
	cur := head
	for cur != nil {
		next := findNext(cur.Next)
		cur.Next = next
		cur = next
	}
	return head
}

func arrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val: arr[0],
		Next: nil,
	}

	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val: arr[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v,", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	arr := []int{}
	l := deleteDuplicates(arrToList(arr))
	printList(l)
}
