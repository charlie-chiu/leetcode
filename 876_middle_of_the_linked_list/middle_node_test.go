package middle_node

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
		{NewLinkedList(1, 2, 3, 4, 5), NewLinkedList(3, 4, 5)},
		{NewLinkedList(1, 2, 3, 4, 5, 6), NewLinkedList(4, 5, 6)},
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
