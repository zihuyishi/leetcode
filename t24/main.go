package main

import "fmt"

func printList(head ListNode) {
	cur := &head
	for cur != nil {
		fmt.Printf("%v, ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{
		Val: 0,
		Next: nil,
	}

	cur := head

	for i := 1; i < 10000; i++ {
		node := &ListNode{
			Val: i,
			Next: nil,
		}

		cur.Next = node
		cur = node
	}

	printList(*head)

	head = swapPairs(head)

	printList(*head)
}
