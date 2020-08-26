package odd_even_linked_list

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	head, want *ListNode
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList([]int{1, 2, 1, 2}), NewLinkedList([]int{1, 1, 2, 2})},
		{NewLinkedList([]int{1, 2, 3, 4, 5}), NewLinkedList([]int{1, 3, 5, 2, 4})},
		{NewLinkedList([]int{2, 1, 3, 5, 6, 4, 7}), NewLinkedList([]int{2, 3, 6, 7, 1, 5, 4})},
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
