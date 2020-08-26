package odd_even_linked_list

import (
	"log"

	. "leetcode"
)

// 12ms / 3.7MB
func oddEvenList(head *ListNode) *ListNode {
	// The length of the linked list is between [0, 10^4].
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	// start from 2th node
	isEvenNode := true
	currNode := head.Next

	lastOdd := head
	firstEven := head.Next

	var oldNext *ListNode

	//log.Println(head)
	for currNode != nil {
		oldNext = currNode.Next
		if isEvenNode {
			log.Println("even node")
			if currNode.Next != nil {
				//skip odd node
				currNode.Next = currNode.Next.Next
			}
		} else {
			//log.Println("odd node")
			lastOdd.Next = currNode
			currNode.Next = firstEven
			lastOdd = currNode
		}

		isEvenNode = !isEvenNode
		currNode = oldNext
		//log.Println(head)
	}

	return head
}
