package main

func main() {
}

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := getLen(l1)
	len2 := getLen(l2)

	var longer, shoter, pre *ListNode

	if len1 >= len2 {
		longer = l1
		shoter = l2
	} else {
		longer = l2
		shoter = l1
	}

	ans := longer
	f := 0

	for shoter != nil {
		sum := longer.Val + shoter.Val + f
		f = 0

		if sum >= 10 {
			sum -= 10
			f++
		}

		longer.Val = sum

		pre = longer
		longer = longer.Next
		shoter = shoter.Next
	}

	if f != 0 {
		if len1 == len2 {
			tnode := &ListNode{
				Val:  1,
				Next: nil,
			}
			pre.Next = tnode
			return ans
		}
		if longer.Val+f < 10 {
			longer.Val = longer.Val + f
			return ans
		}

		pp := longer
		for longer != nil {
			s := longer.Val + f
			if s < 10 {
				longer.Val = s
				return ans
			} else {
				longer.Val = 0
			}
			longer = longer.Next
		}

		for pp.Next != nil {
			pp = pp.Next
		}

		tnode := &ListNode{
			Val:  1,
			Next: nil,
		}
		pp.Next = tnode
	}

	return ans
}

func getLen(l *ListNode) int {
	len := 0
	for l != nil {
		len++
		l = l.Next
	}

	return len
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
