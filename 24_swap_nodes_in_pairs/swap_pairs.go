package swap_pairs

import (
	. "leetcode"
)

// 0ms / 2.1MB
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	head = head.Next //next will swap to head

	for fakeHead.Next != nil && fakeHead.Next.Next != nil {
		first, second := fakeHead.Next, fakeHead.Next.Next
		first.Next = second.Next
		second.Next = first
		fakeHead.Next = second

		fakeHead = first // first at second position now
	}

	return head
}
