package reverse_list

import (
	"fmt"
	"strings"
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

func (node *ListNode) Dump() {
	for node != nil {
		fmt.Printf("node ptr: %p\n", node)
		fmt.Printf("     val: %v\n", node.Val)
		fmt.Printf("    next: %p\n", node.Next)
		node = node.Next
	}

	fmt.Println("### end of node dump ###")
	fmt.Println("")
}
