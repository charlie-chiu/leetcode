package parition_list

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input  *ListNode
	x      int
	output *ListNode
}

func TestPartition(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList(3, 3, 3), 5, NewLinkedList(3, 3, 3)},
		{NewLinkedList(3, 3, 3), 1, NewLinkedList(3, 3, 3)},
		{NewLinkedList(1, 4, 3, 2, 5, 2), 3, NewLinkedList(1, 2, 2, 4, 3, 5)},
	}

	for _, testCase := range TestCases {
		got := partition(testCase.input, testCase.x)

		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("wrong answer at input %v", testCase.input)
			t.Logf("want %v", testCase.output)
			t.Logf(" got %v", got)
		}
	}
}
