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
		{NewLinkedList(1, 2, 3), NewLinkedList(3, 2, 1)},
		{NewLinkedList(1, 2, 3, 4, 5), NewLinkedList(5, 4, 3, 2, 1)},
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
