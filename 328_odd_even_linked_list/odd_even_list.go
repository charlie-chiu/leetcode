package odd_even_linked_list

import (
	"fmt"
	"strings"
)

func oddEvenList(head *ListNode) *ListNode {
	return head
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
