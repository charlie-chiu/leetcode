package parition_list

import (
	. "leetcode"
)

// 0ms
// has solution but failed to implement myself :(
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyBeforeHead := &ListNode{}
	dummyAfterHead := &ListNode{}
	before := dummyBeforeHead
	after := dummyAfterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}

		head = head.Next

	}

	before.Next = dummyAfterHead.Next
	after.Next = nil
	return dummyBeforeHead.Next
}
