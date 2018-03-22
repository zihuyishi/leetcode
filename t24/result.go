package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur, second := head, head.Next

	if second == nil {
		return cur
	}

	next := second.Next

	second.Next, cur.Next = cur, swapPairs(next)

	return second
}