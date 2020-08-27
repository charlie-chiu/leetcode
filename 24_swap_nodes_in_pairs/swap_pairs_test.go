package swap_pairs

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input, output *ListNode
}

func TestSwapPairs(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList(), NewLinkedList()},
		{NewLinkedList(1), NewLinkedList(1)},
		{NewLinkedList(1, 2), NewLinkedList(2, 1)},
		{NewLinkedList(1, 2, 3, 4), NewLinkedList(2, 1, 4, 3)},
	}

	for _, testCase := range TestCases {
		got := swapPairs(testCase.input)

		if !reflect.DeepEqual(testCase.output, got) {
			t.Errorf("wrong output")
			t.Logf("want %v", testCase.output)
			t.Logf(" got %v", got)
		}
	}
}
