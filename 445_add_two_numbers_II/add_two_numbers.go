package add_two_numbers

import (
	"container/list"
	"log"

	. "leetcode"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := toStack(l1)
	s2 := toStack(l2)

	log.Println(s1)
	log.Println(s2)

	var value1, value2, carry int
	var lastNode *ListNode
	for s1.len() != 0 || s2.len() != 0 || carry != 0 {
		if s1.len() == 0 {
			value1 = 0
		} else {
			value1 = s1.pop()
		}

		if s2.len() == 0 {
			value2 = 0
		} else {
			value2 = s2.pop()
		}

		lastNode = &ListNode{
			Val:  (value1 + value2 + carry) % 10,
			Next: lastNode,
		}

		carry = (value1 + value2 + carry) / 10
	}

	return lastNode
}

func toStack(node *ListNode) *intStack {
	s := &intStack{}
	for node != nil {
		s.push(node.Val)
		node = node.Next
	}

	return s
}

type intStack struct {
	s list.List
}

func (s *intStack) push(val int) {
	s.s.PushBack(val)
}

func (s *intStack) pop() int {
	if s.len() > 0 {
		element := s.s.Back()
		s.s.Remove(element)
		return element.Value.(int)
	}
	panic("pop from empty stack")
}

func (s intStack) len() int {
	return s.s.Len()
}
