package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	n1, n2 := head, head.Next

	for n1 != nil && n2 != nil && n2.Next != nil {
		if n1 == n2 {
			return true
		}
		n1 = n1.Next
		n2 = n2.Next.Next
	}

	return false
}
