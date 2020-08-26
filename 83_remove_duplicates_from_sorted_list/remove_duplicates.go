package odd_even_linked_list

import (
	. "leetcode"
)

func deleteDuplicates(head *ListNode) *ListNode {
	unique := head
	for unique != nil {
		for unique.Next != nil && unique.Val == unique.Next.Val {
			unique.Next = unique.Next.Next
		}
		unique = unique.Next
	}

	return head
}
