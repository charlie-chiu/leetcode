package linked_list_cycle

import (
	"fmt"
	"log"
	"strings"
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	sb := strings.Builder{}

	for node != nil {
		sb.WriteString(fmt.Sprintf("%v -> ", node.Val))
		node = node.Next
	}
	sb.WriteString("nil")
	return sb.String()
}
