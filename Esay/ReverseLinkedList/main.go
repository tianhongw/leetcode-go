package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead := new(ListNode)

	for head != nil {
		next := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = next
	}

	return newHead.Next
}
