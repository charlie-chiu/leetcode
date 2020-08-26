package reverse_list

import (
	. "leetcode"
)

// first try, O (n)
// 0ms / 2.6 MB
func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		head = nextNode
	}

	return prevNode
}
