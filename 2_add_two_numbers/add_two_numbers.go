package add_two_numbers

import (
	. "leetcode"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var valueL1, valueL2, carry int
	var resHead = &ListNode{}
	var result = resHead
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			valueL1 = 0
		} else {
			valueL1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			valueL2 = 0
		} else {
			valueL2 = l2.Val
			l2 = l2.Next
		}

		result.Next = &ListNode{
			Val: (valueL1 + valueL2 + carry) % 10,
		}
		result = result.Next
		carry = (valueL1 + valueL2 + carry) / 10
	}

	return resHead.Next
}
