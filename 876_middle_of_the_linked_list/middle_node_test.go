package middle_node

import (
	"reflect"
	"testing"
)

type TestCase struct {
	head, want *ListNode
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{newLinkedList([]int{1, 2, 3, 4, 5}), newLinkedList([]int{3, 4, 5})},
		{newLinkedList([]int{1, 2, 3, 4, 5, 6}), newLinkedList([]int{4, 5, 6})},
	}

	for _, testCase := range TestCases {
		got := middleNode(testCase.head)

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
