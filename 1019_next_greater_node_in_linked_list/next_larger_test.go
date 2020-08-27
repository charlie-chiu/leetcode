package next_larger

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	head *ListNode
	want []int
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList(2, 1, 5), []int{5, 5, 0}},
		{NewLinkedList(2, 7, 4, 3, 5), []int{7, 0, 5, 5, 0}},
		{NewLinkedList(1, 7, 5, 1, 9, 2, 5, 1), []int{7, 9, 9, 9, 0, 5, 0, 0}},
		{NewLinkedList(9, 7, 6, 7, 6, 9), []int{0, 9, 7, 9, 9, 0}},
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
