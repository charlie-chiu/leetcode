package linked_list_cycle

import (
	"log"

	. "leetcode"
)

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var slow, fast = head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func dumpCyclicLst(node *ListNode) {
	for i := 0; i < 10; i++ {
		log.Println(node.Val)
		log.Println("----")
		node = node.Next
	}
}
