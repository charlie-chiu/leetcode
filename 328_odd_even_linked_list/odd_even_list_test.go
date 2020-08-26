package odd_even_linked_list

import (
	"reflect"
	"testing"
)

type TestCase struct {
	head, want *ListNode
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{newLinkedList([]int{1, 2, 1, 2}), newLinkedList([]int{1, 1, 2, 2})},
		{newLinkedList([]int{1, 2, 3, 4, 5}), newLinkedList([]int{1, 3, 5, 2, 4})},
		{newLinkedList([]int{2, 1, 3, 5, 6, 4, 7}), newLinkedList([]int{2, 3, 6, 7, 1, 5, 4})},
	}

	for _, testCase := range TestCases {
		got := oddEvenList(testCase.head)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("wrong output")
			t.Logf("want %v", testCase.want)
			t.Logf(" got %v", got)
		}
	}
}

func newLinkedList(n []int) *ListNode {
	if len(n) == 1 {
		return &ListNode{
			Val:  n[0],
			Next: nil,
		}
	}

	return &ListNode{
		Val:  n[0],
		Next: newLinkedList(n[1:]),
	}
}
