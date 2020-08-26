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
		{NewLinkedList([]int{1, 1, 2}), NewLinkedList([]int{1, 2})},
		{NewLinkedList([]int{1, 1, 2, 3, 3}), NewLinkedList([]int{1, 2, 3})},
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
