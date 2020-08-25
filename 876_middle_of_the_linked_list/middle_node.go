package middle_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head.Next

	for fast != nil {
		slow = slow.Next

		if fast.Next == nil || fast.Next.Next == nil {
			fast = nil
		} else {
			fast = fast.Next.Next
		}
	}

	return slow
}

// 0ms / 2.1 MB
// brute force
func middleNodeBrute(head *ListNode) *ListNode {
	// brute force
	node := head
	slice := []*ListNode{}

	for node != nil {
		slice = append(slice, node)

		node = node.Next
	}

	return slice[len(slice)/2]
}
