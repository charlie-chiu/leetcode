package reverse_list

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
		{NewLinkedList([]int{1, 2, 3}), NewLinkedList([]int{3, 2, 1})},
		{NewLinkedList([]int{1, 2, 3, 4, 5}), NewLinkedList([]int{5, 4, 3, 2, 1})},
	}

	for _, testCase := range TestCases {
		got := reverseList(testCase.head)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("wrong output")
			t.Logf("want %v", testCase.want)
			t.Logf(" got %v", got)
		}
	}
}
