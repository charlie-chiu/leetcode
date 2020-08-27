package next_larger

import (
	"log"

	. "leetcode"
)

// first try, using slow-fast pointer
// accepted but slow
func nextLargerNodes(head *ListNode) []int {
	log.Printf("list: %v", head)
	var slowPtr, fastPtr *ListNode = head, head
	var largerValue = []int{}

	for slowPtr != nil {
		// move fastPtr until fast.val greater than slow.val or fast is nil
		for fastPtr != nil && !(fastPtr.Val > slowPtr.Val) {
			fastPtr = fastPtr.Next
		}

		for !(slowPtr == fastPtr) {
			// no larger value
			if fastPtr == nil {
				largerValue = append(largerValue, 0)
				fastPtr = slowPtr.Next
			} else {
				largerValue = append(largerValue, fastPtr.Val)
			}

			if slowPtr.Next != nil && slowPtr.Next.Val < slowPtr.Val {
				fastPtr = slowPtr.Next
			}

			slowPtr = slowPtr.Next

			//log.Println(largerValue)
			//log.Printf("slow: %v", slowPtr)
			//log.Printf("fast: %v", fastPtr)
			//log.Println("")
		}

	}

	return largerValue
}
