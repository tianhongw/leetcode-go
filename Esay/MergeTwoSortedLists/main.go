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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	ans := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			newHead.Next = l1
			l1 = l1.Next
		} else {
			newHead.Next = l2
			l2 = l2.Next
		}
		newHead = newHead.Next
	}

	if l1 != nil {
		newHead.Next = l1
	}

	if l2 != nil {
		newHead.Next = l2
	}

	return ans.Next
}
