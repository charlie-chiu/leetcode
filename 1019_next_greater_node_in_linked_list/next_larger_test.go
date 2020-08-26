package next_larger

import (
	"reflect"
	"testing"
)

type TestCase struct {
	head *ListNode
	want []int
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		//{newLinkedList([]int{2, 1, 5}), []int{5, 5, 0}},
		//{newLinkedList([]int{2, 7, 4, 3, 5}), []int{7, 0, 5, 5, 0}},
		//{newLinkedList([]int{1, 7, 5, 1, 9, 2, 5, 1}), []int{7, 9, 9, 9, 0, 5, 0, 0}},
		{newLinkedList([]int{9, 7, 6, 7, 6, 9}), []int{0, 9, 7, 9, 9, 0}},
	}

	for _, testCase := range TestCases {
		got := nextLargerNodes(testCase.head)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("wrong answer on list %v", testCase.head)
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
