package leetcode

import (
	"fmt"
	"log"
	"strings"
)

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

func NewLinkedList(values ...int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, value := range values {
		cur.Next = &ListNode{
			Val: value,
		}
		cur = cur.Next
	}

	return dummyHead.Next
}

func DumpCyclicLst(node *ListNode) {
	var count int
	sb := strings.Builder{}

	for node != nil && count < 10 {
		sb.WriteString(fmt.Sprintf("%v -> ", node.Val))
		node = node.Next
		count++
	}
	sb.WriteString("...")
	log.Println(sb.String())
}
