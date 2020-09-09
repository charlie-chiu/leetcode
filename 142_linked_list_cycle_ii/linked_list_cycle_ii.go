package linked_list_cycle_ii

import (
	. "leetcode"
)

//try brute force
// and
//Floyd's algorithm

func detectCycle(head *ListNode) *ListNode {
	return head
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var slow, fast = head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
