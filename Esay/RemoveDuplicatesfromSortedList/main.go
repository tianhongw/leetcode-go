package main

import "sort"

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

func deleteDuplicates(head *ListNode) *ListNode {
	m := getMap(head)
	sortedK := getSortedK(m)

	newHead := new(ListNode)
	ans := newHead

	for _, v := range sortedK {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		newHead.Next = node
		newHead = node
	}

	return ans.Next
}

func getMap(head *ListNode) map[int]bool {
	m := make(map[int]bool)
	for head != nil {
		m[head.Val] = true
		head = head.Next
	}

	return m
}

func getSortedK(m map[int]bool) []int {
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}

	sort.Ints(ret)
	return ret
}
