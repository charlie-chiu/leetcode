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
		//{newLinkedList([]int{1,1,2}), newLinkedList([]int{1,2})},
		{newLinkedList([]int{1, 1, 2, 3, 3}), newLinkedList([]int{1, 2, 3})},
	}

	for _, testCase := range TestCases {
		got := deleteDuplicates(testCase.head)

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
