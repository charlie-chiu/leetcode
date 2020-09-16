package merge_two_sorted_lists

import . "leetcode"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// first try, 0ms / 2.5 MB (99.77%)
	var answer = &ListNode{}
	var answerHead = answer

	for !(l1 == nil && l2 == nil) {
		if l1 == nil {
			answer.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			answer.Next = l1
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				answer.Next = l1
				l1 = l1.Next
			} else {
				answer.Next = l2
				l2 = l2.Next
			}
		}

		answer = answer.Next
	}

	return answerHead.Next
}
